package handlers

import (
	"net/http"
	"uhs/internal/models"
	"uhs/internal/responses"
	"uhs/internal/services"
	"uhs/internal/types"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func GetApplications(applicationOps types.ApplicationOps) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		apps, err := applicationOps.GetEntities()
		if err != nil {
			c.Logger().Error("Failed to get all applications " + err.Error())
			return c.JSON(http.StatusNotFound, &responses.DefaultResponse{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "Applications not found",
			})
		}
		return c.JSON(http.StatusOK, apps)
	}
}

func SubmitApplication(applicationOps types.ApplicationOps) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		var application models.Application
		posId := c.Param("posId")
		if len(posId) == 0 {
			c.Logger().Error("No position id provided for application ")
			return c.JSON(http.StatusInternalServerError, &responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: "invalid position id",
			})
		}
		// get userId from token
		user, err := services.ExtractToken(c)
		if err != nil {
			c.Logger().Error("Failed to retrieve token from request body " + err.Error())
			return c.JSON(http.StatusUnauthorized, &responses.DefaultResponse{
				Status:  http.StatusUnauthorized,
				Success: false,
				Message: "Invalid or Missing JWT",
			})
		}
		// bind required values
		application.Id = uuid.NewString()
		application.UserId = user.UserId
		application.PositionId = posId

		err = applicationOps.CreateEntity(&application)
		if err != nil {
			c.Logger().Error("Failed to create application " + err.Error())
			return c.JSON(http.StatusInternalServerError, &responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: "Failed to create application",
			})
		}
		return c.JSON(http.StatusOK, application)
	}
}
