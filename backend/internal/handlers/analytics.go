package handlers

import (
	"net/http"
	"uhs/internal/responses"
	"uhs/internal/services"
	"uhs/internal/types"
	"uhs/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func SaveAnalytics(analyticsOps types.AnalyticsOps, pdfOps types.PDFOps, aiOps types.AIOps, cvOps types.CvOps) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		// extract token
		token, err := echo.ContextGet[*jwt.Token](c, "user")
		if err != nil {
			c.Logger().Error("Failed to retrieve token from request body " + err.Error())
			return c.JSON(http.StatusUnauthorized, &responses.DefaultResponse{
				Status:  http.StatusUnauthorized,
				Success: false,
				Message: "Invalid or Missing JWT",
			})
		}
		user := token.Claims.(*services.CustomJWTClaims)
		// find CV
		cv := cvOps.GetEntitiesWhere("user_id = ?", user.UserId)
		if len(cv) == 0 {
			c.Logger().Error("No such user found ")
			return c.JSON(http.StatusUnauthorized, &responses.DefaultResponse{
				Status:  http.StatusUnauthorized,
				Success: false,
				Message: "Invalid or Missing JWT",
			})
		}
		// initiate
		text, err := pdfOps.ExtractText(c, cv[0].FileUrl)
		if err != nil {
			c.Logger().Error("Failed to extrac pdf text" + err.Error())
			return c.JSON(http.StatusInternalServerError, &responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: "Failed to process pdf",
			})
		}

		// ai
		analytics, err := aiOps.ProcessCV(c, utils.GetPrompt("The Job is about Lab Engineer", text))
		// bind omitted values
		analytics.Id = uuid.NewString()
		analytics.UserId = user.UserId

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
