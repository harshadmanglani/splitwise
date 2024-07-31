package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
	"time"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar = logger.Sugar()
	defer logger.Sync()
}

type JwtManager struct {
	secret    string
	algorithm Algorithm
}

func NewJwtManager(secret string, algorithm Algorithm) *JwtManager {
	jm := &JwtManager{
		secret:    secret,
		algorithm: algorithm,
	}
	return jm
}

func (jm *JwtManager) Generate(claims Claims) string {
	headers := headers{
		Alg: jm.algorithm,
		Typ: "JWT",
	}
	base64Headers := base64Encode(headers)
	base64Claims := base64Encode(claims)
	signature := jm.generateJwtSignature(strings.Join([]string{base64Headers, base64Claims}, "."))
	return strings.Join([]string{base64Headers, base64Claims, signature}, ".")
}

func (jm *JwtManager) generateJwtSignature(message string) string {
	mac := hmac.New(sha256.New, []byte(jm.secret))
	mac.Write([]byte(message))
	signature := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature)
}

func (jm *JwtManager) VerifyAndReturnClaims(token string) (Claims, AuthError) {
	splitToken := strings.Split(token, ".")
	base64Headers := splitToken[0]
	base64Claims := splitToken[1]
	base64Signature := splitToken[2]

	generatedBase64Signature := jm.generateJwtSignature(strings.Join([]string{base64Headers, base64Claims}, "."))
	var claims Claims
	if subtle.ConstantTimeCompare([]byte(base64Signature), []byte(generatedBase64Signature)) == 1 {
		jsonClaims, err := base64.StdEncoding.DecodeString(base64Claims)
		if err != nil {
			sugar.Error("An error occurred in decrypting the token!")
			panic(err)
		}
		json.Unmarshal(jsonClaims, &claims)
		if time.Now().After(claims.Expiry) {
			return Claims{}, TOKEN_EXPIRED
		}
		return claims, NO_ERROR
	}
	return Claims{}, INVALID_SIGNATURE
}

func base64Encode(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(jsonData)
}
