package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type JwtGenerator struct {
	secret    string
	algorithm Algorithm
}

func NewJwtGenerator(secret string, algorithm Algorithm) *JwtGenerator {
	generator := &JwtGenerator{
		secret:    secret,
		algorithm: algorithm,
	}
	return generator
}

func (generator *JwtGenerator) GenerateJwt(claims Claims) string {
	headers := headers{
		Alg: generator.algorithm,
		Typ: "JWT",
	}
	base64Headers := base64Encode(headers)
	base64Claims := base64Encode(claims)
	signature := generator.generateJwtSignature(strings.Join([]string{base64Headers, base64Claims}, "."))
	return strings.Join([]string{base64Headers, base64Claims, signature}, ".")
}

func (generator *JwtGenerator) generateJwtSignature(message string) string {
	mac := hmac.New(sha256.New, []byte(generator.secret))
	mac.Write([]byte(message))
	signature := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature)
}

func (generator *JwtGenerator) VerifyAndReturnClaims(token string) (Claims, AuthError) {
	splitToken := strings.Split(token, ".")
	base64Headers := splitToken[0]
	base64Claims := splitToken[1]
	base64Signature := splitToken[2]

	generatedBase64Signature := generator.generateJwtSignature(strings.Join([]string{base64Headers, base64Claims}, "."))
	var claims Claims
	if subtle.ConstantTimeCompare([]byte(base64Signature), []byte(generatedBase64Signature)) == 1 {
		jsonClaims, err := base64.StdEncoding.DecodeString(base64Claims)
		if err != nil {
			fmt.Println("An error occurred in decrypting the token!")
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
