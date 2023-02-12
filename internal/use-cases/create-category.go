package use_cases

import (
	"github.com/GabrielBrotas/go-categories-msvc/internal/entities"
	"github.com/GabrielBrotas/go-categories-msvc/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	// TODO: verify if category name already exists
	err = u.repository.Save(category)

	if err != nil {
		return err
	}

	return nil
}
