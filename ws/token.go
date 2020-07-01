package ws

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

// JwtClaims custom claims
type JwtClaims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

// VerifyToken verify token
func verifyToken(secret, tokenStr string) (*JwtClaims, error) {
	var m jwt.MapClaims

	token, err := jwt.ParseWithClaims(tokenStr, &m, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	var claims JwtClaims
	mapstructure.Decode(&m, &claims)

	return &claims, nil
}
