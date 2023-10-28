package jwt

import "time"

type Algorithm int

const (
	HMACSHA256 Algorithm = iota
)

type Claims struct {
	Issuer   string                 `json:"iss"`
	Subject  string                 `json:"sub"`
	Expiry   time.Time              `json:"exp"`
	IssuedAt time.Time              `json:"iat"`
	Custom   map[string]interface{} `json:"custom"`
}

type headers struct {
	Alg Algorithm `json:"alg"`
	Typ string    `json:"typ"`
}

type AuthError int

const (
	TOKEN_EXPIRED     AuthError = iota
	INVALID_SIGNATURE AuthError = iota
	NO_ERROR          AuthError = iota
)
