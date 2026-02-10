package handlers

import (
	"net/http"
	"uhs/internal/models"
	"uhs/internal/responses"
	"uhs/internal/services"
	"uhs/internal/types"
	"uhs/internal/utils"

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

func SubmitApplication(applicationOps types.ApplicationOps, positionOps types.PositionOps, cvOps types.CvOps, pdfOps types.PDFOps, aiOps types.AIOps, analyticsOps types.AnalyticsOps) func(c *echo.Context) error {
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
		// check if position exists for this id
		pos, err := positionOps.GetEntityByID(posId)
		if err != nil {
			c.Logger().Error("Position not found in the position table of database")
			return c.JSON(http.StatusNotFound, &responses.DefaultResponse{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "position not found",
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

		// Save Analytics

		// find CV
		cvList := cvOps.GetEntitiesWhere("user_id = ?", user.UserId)
		if len(cvList) == 0 {
			c.Logger().Error("CV not found in the cv table of database")
			return c.JSON(http.StatusNotFound, &responses.DefaultResponse{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "cv not found",
			})
		}
		cv := cvList[0]

		// initiate

		text, err := pdfOps.ExtractText(c, cv.FileUrl)
		if err != nil {
			c.Logger().Error("Failed to extract pdf text" + err.Error())
			return c.JSON(http.StatusInternalServerError, &responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: "Failed to process pdf",
			})
		}
		// ai
		analytics, err := aiOps.ProcessCV(c, utils.GetPrompt(pos.Description, text))
		if err != nil {
			c.Logger().Error("Failed to connect to AI model for processing..")
			return c.JSON(http.StatusInternalServerError, &responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: "Failed connecting to Groq Client",
			})
		}

		// bind omitted values
		analytics.Id = uuid.NewString()
		analytics.UserId = user.UserId
		analytics.PositionId = posId
		// save to db
		err = analyticsOps.CreateEntity(&analytics)
		if err != nil {
			c.Logger().Error("Failed to create analytics " + err.Error())
			return c.JSON(http.StatusBadRequest, &responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "Failed to create analytics",
			})
		}

		// Now Application is complete just save it

		// bind required values
		application.Id = uuid.NewString()
		application.UserId = user.UserId
		application.PositionId = posId

		err = applicationOps.CreateEntity(&application)
		if err != nil {
			c.Logger().Error("Failed to create application " + err.Error())
			return c.JSON(http.StatusBadRequest, &responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "Application already exists!",
			})
		}

		return c.JSON(http.StatusOK, responses.DefaultResponse{
			Status:  http.StatusOK,
			Success: true,
			Message: "Thanks for Applying",
		})
	}
}
