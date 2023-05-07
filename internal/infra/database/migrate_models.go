package database

import (
	c_models "github.com/GabrielBrotas/go-categories-msvc/internal/categories/models"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&c_models.Category{})
}
