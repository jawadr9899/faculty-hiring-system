package services

import (
	"uhs/internal/config"
	"uhs/internal/models"
	"uhs/internal/types/common"

	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DatabaseOperations[T any] interface {
	CreateEntity(entity *T) error
	GetEntities() ([]T, error)
	GetEntityByID(id string) (T, error)
	GetEntitiesWhere(query string, placeholders ...any) []T
}

type Database[T any] struct {
	Db *gorm.DB
}

func SetupDB(app *echo.Echo, cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		app.Logger.Error("Failed to connect to database")
		return nil, err
	}
	// migrations
	db.AutoMigrate(&models.User{}, &models.Analytics{}, &models.Application{}, &models.Position{}, &models.Cv{})
	// set admin

	if err := saveAdmin(app, db, cfg); err != nil {
		app.Logger.Error("Failed to create admin user")
		return nil, err
	}
	return db, nil
}

func NewDatabaseService[T any](db *gorm.DB) DatabaseOperations[T] {
	return &Database[T]{
		Db: db,
	}
}

// Sets Admin
func saveAdmin(app *echo.Echo, db *gorm.DB, cfg *config.Config) error {
	var adminCount int64
	db.Model(&models.User{}).Where("role = ?", "admin").Count(&adminCount)

	if adminCount > 0 {
		app.Logger.Info("Admin already exists with count ")
		return nil
	}

	adminUUID := uuid.NewString()
	hash, err := bcrypt.GenerateFromPassword([]byte(cfg.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		app.Logger.Error("Failed to hash the password of admin")
		return err
	}

	admin := &models.User{
		Id:       adminUUID,
		Name:     "admin",
		Email:    cfg.AdminEmail,
		Password: string(hash),
		Role:     common.AdminRole,
	}

	err = db.Create(&admin).Error
	if err != nil {
		app.Logger.Error("Failed to save admin in database")
		return err
	}
	return nil
}

// Generic functions
func (db *Database[T]) CreateEntity(entity *T) error {
	return db.Db.Create(entity).Error
}

func (db *Database[T]) GetEntities() ([]T, error) {
	var entities []T
	err := db.Db.Find(&entities).Error
	return entities, err
}

func (db *Database[T]) GetEntityByID(id string) (T, error) {
	var entity T
	err := db.Db.First(&entity, "id = ?", id).Error
	return entity, err
}

func (db *Database[T]) GetEntitiesWhere(query string, placeholders ...any) []T {
	var entities []T
	db.Db.Where(query, placeholders...).Find(&entities)
	return entities
}
