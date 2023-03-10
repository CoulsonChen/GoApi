package services

import (
	"time"

	"github.com/CoulsonChen/GoApi/pkg/config"
	"github.com/CoulsonChen/GoApi/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

type IJwtService interface {
	GenToken(account string) (string, error)
	ParseToken(tokenString string) (*models.MyClaims, error)
}

type JwtService struct {
	jwtconfig config.JwtConfig
}

func JwtServiceProvider(jwtconfig config.JwtConfig) *JwtService {
	return &JwtService{
		jwtconfig: jwtconfig,
	}
}

func (j *JwtService) GenToken(account string) (string, error) {
	c := models.MyClaims{
		account,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(j.jwtconfig.JwtTokenExpireDuration)).Unix(),
			Issuer:    "Coulson",
		},
	}
	// Choose specific algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Choose specific Signature
	return token.SignedString([]byte(j.jwtconfig.JwtSecretKey))
}

func (j *JwtService) ParseToken(tokenString string) (*models.MyClaims, error) {
	jwt, err := jwt.ParseWithClaims(tokenString, &models.MyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(j.jwtconfig.JwtSecretKey), nil
	})
	if err == nil && jwt != nil {
		if claim, ok := jwt.Claims.(*models.MyClaims); ok && jwt.Valid {
			return claim, nil
		}
	}
	// panic(err)
	return nil, err
}
