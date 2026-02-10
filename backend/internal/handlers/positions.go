package handlers

import (
	"net/http"
	"uhs/internal/models"
	"uhs/internal/responses"
	"uhs/internal/types"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func GetPositions(positionOps types.PositionOps) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		positions, err := positionOps.GetEntities()
		if err != nil {
			c.Logger().Error("Failed to get all positions " + err.Error())
			return c.JSON(http.StatusUnauthorized, &responses.DefaultResponse{
				Status:  http.StatusUnauthorized,
				Success: false,
				Message: "Invalid or Missing JWT",
			})
		}
		return c.JSON(http.StatusOK, positions)
	}
}

func AddPosition(positionOps types.PositionOps) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		var position models.Position
		err := echo.BindBody(c, &position)
		if err != nil {
			c.Logger().Error("Failed to read request body " + err.Error())
			return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "Failed to read request body",
			})
		}
		// bind values
		position.Id = uuid.NewString()
		// add to db
		err = positionOps.CreateEntity(&position)
		if err != nil {
			c.Logger().Error("Failed to add position to database " + err.Error())
			return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: "Failed to add position entity to database",
			})
		}
		return c.JSON(http.StatusOK, responses.DefaultResponse{
			Status:  http.StatusOK,
			Success: true,
			Message: position,
		})

	}
}
