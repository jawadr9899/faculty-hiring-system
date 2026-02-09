package handlers

import (
	"net/http"
	"time"
	"uhs/internal/config"
	"uhs/internal/models"
	"uhs/internal/responses"
	"uhs/internal/services"
	"uhs/internal/types"
	"uhs/internal/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(userOps types.UserOps) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		users, err := userOps.GetEntities()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: err,
			})
		}
		return c.JSON(http.StatusOK, users)
	}

}

func Signup(userOps types.UserOps, cvOps types.CvOps, cfg *config.Config) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		var user models.User
		user.Id = uuid.NewString()
		user.Name = c.FormValue("name")
		user.Email = c.FormValue("email")
		user.Password = c.FormValue("password")
		
		// get cv file from multipart
		file, err := c.FormFile("file")
		if err != nil {
			c.Logger().Error("Failed to retrieve file from multipart " + err.Error())
			return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: err,
			})
		}
		// save it to server
		savedPath, err := utils.SaveFileInServer(c, file, cfg)
		if err != nil {
			c.Logger().Error("Failed to save file to server " + err.Error())
			return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: err,
			})
		}
		// Save file info to file database
		err = cvOps.CreateEntity(&models.Cv{
			Id:       uuid.NewString(),
			UserId:   user.Id,
			FileName: file.Filename,
			FileUrl:  savedPath,
		})
		if err != nil {
			c.Logger().Error("Failed to save file info to file database " + err.Error())
			return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: err,
			})
		}

		// check if user already exists in database
		usersList := userOps.GetEntitiesWhere("email = ?", user.Email)
		if len(usersList) > 0 {
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
		err = userOps.CreateEntity(&user)
		if err != nil {
			c.Logger().Error("Database failed to put user")
			return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, user)
	}
}

func Login(userOps types.UserOps) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		var user models.User
		err := echo.BindBody(c, &user)
		if err != nil {
			c.Logger().Error("No body found in request")
			return err
		}
		// check if user exists or not in database
		usersList := userOps.GetEntitiesWhere("email = ?", user.Email)
		if len(usersList) == 0 {
			c.Logger().Error("User doesn't exist")
			return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "Invalid credentials",
			})
		}
		// check the hash
		err = bcrypt.CompareHashAndPassword([]byte(usersList[0].Password), []byte(user.Password))
		if err != nil {
			c.Logger().Error("Invalid credentials")
			return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "Invalid credentials",
			})
		}
		// generate jwt token
		claims := services.NewCustomClaims(usersList[0].Id, usersList[0].Email, time.Now().Add(time.Minute*15))
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
