package routes

import (
	"uhs/internal/handlers"
	"uhs/internal/repository"

	"github.com/labstack/echo/v5"
)

func SetupRoutes(api *echo.Group, db *repository.Sqlite) {
	api.GET("/user", handlers.GetUsers(db))
	api.POST("/user",handlers.PostUser(db))

}
