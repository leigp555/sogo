package token

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}
