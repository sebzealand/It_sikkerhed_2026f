package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("hemmelig-nøgle123")

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iss":   "my-auth-server",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return fmt.Errorf("Invalid token!")
	}

	return nil
}
