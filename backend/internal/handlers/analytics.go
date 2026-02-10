package handlers

import (
	"net/http"
	"uhs/internal/responses"
	"uhs/internal/services"
	"uhs/internal/types"
	"uhs/internal/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func SaveAnalytics(analyticsOps types.AnalyticsOps, pdfOps types.PDFOps, aiOps types.AIOps, cvOps types.CvOps, positionOps types.PositionOps) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		// get token out of this shi
		user, err := services.ExtractToken(c)
		if err != nil {
			c.Logger().Error("Failed to retrieve token from request body " + err.Error())
			return c.JSON(http.StatusUnauthorized, &responses.DefaultResponse{
				Status:  http.StatusUnauthorized,
				Success: false,
				Message: "Invalid or Missing JWT",
			})
		}
		// get the position for which analytics are required
		posId := c.Param("posId")
		if len(posId) == 0 {
			c.Logger().Error("Position Id not found how do I retrieve it from DB ???")
			return c.JSON(http.StatusBadRequest, &responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "invalid position id",
			})
		}
		// get the position for which user is applying in db
		pos, err := positionOps.GetEntityByID(posId)
		if err != nil {
			c.Logger().Error("Position not found in the position table of database")
			return c.JSON(http.StatusNotFound, &responses.DefaultResponse{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "position not found",
			})
		}
		// find CV
		cv, err := cvOps.GetEntityByID(user.UserId)
		if err != nil {
			c.Logger().Error("CV not found in the cv table of database")
			return c.JSON(http.StatusNotFound, &responses.DefaultResponse{
				Status:  http.StatusNotFound,
				Success: false,
				Message: "cv not found",
			})
		}
		// initiate
		text, err := pdfOps.ExtractText(c, cv.FileUrl)
		if err != nil {
			c.Logger().Error("Failed to extrac pdf text" + err.Error())
			return c.JSON(http.StatusInternalServerError, &responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: "Failed to process pdf",
			})
		}
		// ai
		analytics, err := aiOps.ProcessCV(c, utils.GetPrompt(pos.Description, text))
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
		return c.JSON(http.StatusOK, &responses.DefaultResponse{
			Status:  200,
			Success: true,
			Message: map[string]any{"analytics": analytics},
		})
	}
}
