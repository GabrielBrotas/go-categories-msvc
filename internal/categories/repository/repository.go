package repository

import (
	"errors"

	c_models "github.com/GabrielBrotas/go-categories-msvc/internal/categories/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (r *CategoryRepository) FindById(id uint) (*c_models.Category, error) {
	var category c_models.Category
	result := r.db.First(&category, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &category, nil
}

func (r *CategoryRepository) FindByName(name string) (*c_models.Category, error) {
	var category c_models.Category

	result := r.db.Where("name = ?", name).First(&category)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &category, nil
}

// TODO: Pagination
func (r *CategoryRepository) FindAll() ([]*c_models.Category, error) {
	var categories []*c_models.Category
	result := r.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (r *CategoryRepository) Create(category *c_models.Category) error {
	result := r.db.Create(category)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CategoryRepository) Update(category *c_models.Category) error {
	result := r.db.Save(category)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CategoryRepository) Delete(id uint) error {
	var category c_models.Category
	result := r.db.First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	result = r.db.Delete(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
