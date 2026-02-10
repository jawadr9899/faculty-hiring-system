package services

import (
	"log"
	"time"
	"uhs/internal/config"
	"uhs/internal/types/common"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

type AuthService interface {
	GenerateToken(userId string, userName string) (string, error)
	ExtractToken(c *echo.Context) (*CustomJWTClaims, error)
}

type CustomJWTClaims struct {
	UserId   string      `json:"userid"`
	UserName string      `json:"username"`
	Role     common.Role `json:"role"`
	jwt.RegisteredClaims
}

func NewCustomClaims(userId string, userName string, role common.Role, exp time.Time) *CustomJWTClaims {
	return &CustomJWTClaims{
		UserId:   userId,
		UserName: userName,
		Role:     role,
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

// Extracts Token from a echo Context (*echo.Context)
func ExtractToken(c *echo.Context) (*CustomJWTClaims, error) {
	token, err := echo.ContextGet[*jwt.Token](c, "user")
	if err != nil {
		return nil, err
	}
	user := token.Claims.(*CustomJWTClaims)
	return user, nil
}
