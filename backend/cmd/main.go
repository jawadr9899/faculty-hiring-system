package main

import (
	"log"
	"os"
	"uhs/internal/config"
	"uhs/internal/repository"
	"uhs/internal/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load env file with error\n%s", err)
	}
	cfg := config.LoadConfig()
	app := echo.New()
	app.Use(middleware.RequestLogger())
	api := app.Group("/api")

	// DB Connection
	db, err := repository.New(cfg.SqliteDBPath, app)
	if err != nil {
		app.Logger.Error("Failed to connect to database")
		os.Exit(1)
	}

	// Routes
	routes.SetupRoutes(api, db)

	if err := app.Start(":8080"); err != nil {
		app.Logger.Error("Failed to start server with error ")
	}
}
