package models

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}
