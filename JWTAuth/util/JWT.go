package util

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("haoidhawckla791ahdiw832jlwfow38")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}