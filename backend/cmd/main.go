package main

import (
	"log"
	"uhs/internal/config"
	"uhs/internal/routes"
	"uhs/internal/services"

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
	// Middlewares
	app.Use(middleware.RequestLogger())
	app.Use(middleware.Recover())

	// Groups
	api := app.Group("/api")
	// DB Connection
	db, err := services.SetupDB(app, cfg)
	if err != nil {
		app.Logger.Error("Failed to connect to database")
	}
	// Routes
	routes.SetupRoutes(cfg, api, db)

	if err := app.Start(":8080"); err != nil {
		app.Logger.Error("Failed to start server with error ")
	}
}
