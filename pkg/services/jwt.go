package services

import (
	"errors"
	"time"

	"github.com/CoulsonChen/GoApi/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

type IJwtService interface {
	GenToken(account string) (string, error)
	ParseToken(tokenString string) (*models.MyClaims, error)
}

type JwtService struct {
	JwtSecretKey           string
	JwtTokenExpireDuration int
}

func JwtServiceProvider() *JwtService {
	return &JwtService{
		JwtSecretKey:           "cc2e8a10-2512-4024-8580-84f06109630c",
		JwtTokenExpireDuration: 120,
	}
}

func (j *JwtService) GenToken(account string) (string, error) {
	c := models.MyClaims{
		account,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(j.JwtTokenExpireDuration)).Unix(),
			Issuer:    "Coulson",
		},
	}
	// Choose specific algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Choose specific Signature
	return token.SignedString([]byte(j.JwtSecretKey))
}

func (j *JwtService) ParseToken(tokenString string) (*models.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return j.JwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	// Valid token
	if claims, ok := token.Claims.(*models.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
