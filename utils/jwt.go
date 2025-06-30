package utils

import (
	"KYC/iternals/configs"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	UserID int    `json:"user_id"`
	Login  string `json:"login"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userID int, login string, role string) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		Login:  login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Duration(configs.AppSettings.AuthParams.JwtTtlMinutes) * time.Minute),
			Issuer:    configs.AppSettings.AppParams.ServerName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
