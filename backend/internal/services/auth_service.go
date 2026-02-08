package services

import (
	"log"
	"time"
	"uhs/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	GenerateToken(userId string, userName string) (string, error)
}

type CustomJWTClaims struct {
	UserId   string `json:"userid"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

func NewCustomClaims(userId string, userName string, exp time.Time) *CustomJWTClaims {
	return &CustomJWTClaims{
		UserId:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
}

func (claims *CustomJWTClaims) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encoded_token, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET", "JON_SNOW")))

	if err != nil {
		log.Fatalf("Failed to generate the jwt token %s", err)
		return "", err
	}
	return encoded_token, nil

}
