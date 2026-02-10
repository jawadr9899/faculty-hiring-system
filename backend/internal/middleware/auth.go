package middleware

import (
	"net/http"
	"uhs/internal/config"
	"uhs/internal/responses"
	"uhs/internal/services"
	"uhs/internal/types/common"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

func Authenticate(cfg *config.Config) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(cfg.JWTSecret),
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(services.CustomJWTClaims)
		},
	})
}

func Authorize() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			token, err := services.ExtractToken(c)
			if err != nil {
				c.Logger().Error("Failed to authorize, jwt(token) not found")
				return c.JSON(http.StatusUnauthorized, responses.DefaultResponse{
					Status:  http.StatusUnauthorized,
					Success: false,
					Message: "Malformed or Invalid JWT",
				})
			}
			if token.Role != common.AdminRole {
				c.Logger().Error("Unauthorized, Failed user trying to access admin")
				return c.JSON(http.StatusForbidden, responses.DefaultResponse{
					Status:  http.StatusForbidden,
					Success: false,
					Message: echo.ErrForbidden.Error(),
				})
			}

			return next(c)
		}
	}
}
