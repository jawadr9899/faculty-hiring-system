package routes

import (
	"uhs/internal/config"
	"uhs/internal/handlers"
	"uhs/internal/middleware"
	"uhs/internal/models"
	"uhs/internal/services"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func SetupRoutes(cfg *config.Config, api *echo.Group, db *gorm.DB) {
	// db services
	userServiceOps := services.NewDatabaseService[models.User](db)
	analyticsServiceOps := services.NewDatabaseService[models.Analytics](db)
	cvServiceOps := services.NewDatabaseService[models.Cv](db)
	positionServiceOps := services.NewDatabaseService[models.Position](db)

	// extra services
	pdfServiceOps := services.NewPDFService()
	aiServiceOps := services.NewAIService(cfg.AiApiKey, cfg.AiApiModel)

	// user related routes
	usersApi := api.Group("/user")
	usersApi.GET("/all", handlers.GetUsers(userServiceOps), middleware.Authenticate(cfg))
	usersApi.POST("/signup", handlers.Signup(userServiceOps, cvServiceOps, cfg))
	usersApi.POST("/login", handlers.Login(userServiceOps))
	// protected routes
	usersApi.POST("/analytics", handlers.SaveAnalytics(analyticsServiceOps, pdfServiceOps, aiServiceOps, cvServiceOps), middleware.Authenticate(cfg))

	posApi := api.Group("/position")
	// job posting related routes
	posApi.GET("/all", handlers.GetPositions(positionServiceOps))
	posApi.POST("/post", handlers.AddPosition(positionServiceOps))

}
