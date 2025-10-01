package entity

import "github.com/golang-jwt/jwt/v5"

type Token struct {
	Access  string
	Refresh string
}

type Claims struct {
	Sub  int64
	Role string
	jwt.RegisteredClaims
}
