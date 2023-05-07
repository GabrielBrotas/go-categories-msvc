package use_cases

import (
	"errors"

	c_models "github.com/GabrielBrotas/go-categories-msvc/internal/categories/models"
	c_repository "github.com/GabrielBrotas/go-categories-msvc/internal/categories/repository"
)

type GetCategoryUseCase struct {
	repository *c_repository.CategoryRepository
}

func NewGetCategoryUseCase(repository *c_repository.CategoryRepository) *GetCategoryUseCase {
	return &GetCategoryUseCase{repository}
}

func (u *GetCategoryUseCase) Execute(id uint) (*c_models.Category, error) {
	category, err := u.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, errors.New("category not found")
	}

	return category, nil
}
