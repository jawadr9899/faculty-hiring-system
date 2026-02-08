package handlers

import (
	"fmt"
	"net/http"
	"time"
	"uhs/internal/models"
	"uhs/internal/repository"
	"uhs/internal/responses"
	"uhs/internal/services"
	"uhs/internal/types"

	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(db *repository.Sqlite) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		users, err := db.GetAllUsers()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: err,
			})
		}
		return responses.JsonResponse(c, http.StatusOK, users...)
	}

}

func Signup(db *repository.Sqlite) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		var user models.User
		err := echo.BindBody(c, &user)
		if err != nil {
			c.Logger().Error("Failed to post user")
			return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: err,
			})
		}
		// check if user already exists in database
		_, err = db.GetUserByCol("email", user.Email)
		if err == nil {
			c.Logger().Error("User already exists")
			return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "User already exists",
			})
		}
		// hash password
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.Logger().Error("Failed to hash password")
			return err
		}
		user.Password = string(hash)

		code, err := db.CreateUser(user)
		if err != nil {
			c.Logger().Error("Database failed to put user")
			return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: fmt.Sprintf("%s %d", err.Error(), code),
			})
		}
		return responses.JsonResponse(c, http.StatusOK, user)
	}
}

func Login(db *repository.Sqlite) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		var user types.UserLogin
		err := echo.BindBody(c, &user)
		if err != nil {
			c.Logger().Error("No body found in request")
			return err
		}
		// check if user exists or not in database
		u, err := db.GetUserByCol("email", user.Email)
		if err != nil {
			c.Logger().Error("User doesn't exist")
			return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "Invalid credentials",
			})
		}
		// check the hash
		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
		if err != nil {
			c.Logger().Error("Invalid credentials")
			return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "Invalid credentials",
			})
		}
		// generate jwt token
		claims := services.NewCustomClaims(*u.Id, u.Email, time.Now().Add(time.Minute*15))
		token, err := claims.GenerateToken()

		if err != nil {
			c.Logger().Error("Failed to generated jwt token")
			return err
		}
		return c.JSON(http.StatusOK, &responses.DefaultResponse{
			Status:  http.StatusOK,
			Success: true,
			Message: map[string]string{
				"token": token,
			},
		})

	}
}
