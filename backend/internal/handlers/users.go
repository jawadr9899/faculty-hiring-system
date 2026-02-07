package handlers

import (
	"fmt"
	"net/http"
	"uhs/internal/repository"
	"uhs/internal/responses"
	"uhs/internal/types"

	"github.com/labstack/echo/v5"
)

func GetUsers(db *repository.Sqlite) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		users, err := db.GetAllUsers()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: err,
			})
		}
		return responses.JsonResponse(c, http.StatusOK, users...)
	}

}

func PostUser(db *repository.Sqlite) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		var user types.User
		err := echo.BindBody(c, &user)
		if err != nil {
			c.Logger().Error("Failed to post user")
			return c.JSON(http.StatusBadRequest, &responses.UserResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: err,
			})
		}
		code, err := db.CreateUser(user)
		if err != nil {
			c.Logger().Error("Database failed to put user")
			return c.JSON(http.StatusBadRequest, &responses.UserResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: fmt.Sprintf("%s %d", err.Error(), code),
			})
		}
		return responses.JsonResponse(c, http.StatusOK, user)
	}
}
