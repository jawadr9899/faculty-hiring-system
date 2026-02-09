package services

import (
	"uhs/internal/config"
	"uhs/internal/models"

	"github.com/glebarez/sqlite"
	"github.com/labstack/echo/v5"
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
	return db, nil
}

func NewDatabaseService[T any](db *gorm.DB) DatabaseOperations[T] {
	return &Database[T]{
		Db: db,
	}
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
