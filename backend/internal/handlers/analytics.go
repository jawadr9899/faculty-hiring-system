package handlers

import (
	"net/http"
	"uhs/internal/responses"
	"uhs/internal/types"

	"github.com/labstack/echo/v5"
)

func GetAllAnalytics(analyticsOps types.AnalyticsOps) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		analytics, err := analyticsOps.GetEntities()
		if err != nil {
			c.Logger().Error("Failed to retrieve the analytics from database")
			return c.JSON(http.StatusNotFound, responses.DefaultResponse{
				Status:  http.StatusNotFound,
				Success: false,
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, &responses.DefaultResponse{
			Status:  200,
			Success: true,
			Message: map[string]any{"analytics": analytics},
		})
	}
}
