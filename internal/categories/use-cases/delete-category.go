package use_cases

import (
	c_repository "github.com/GabrielBrotas/go-categories-msvc/internal/categories/repository"
)

type DeleteCategoryUseCase struct {
	repository *c_repository.CategoryRepository
}

func NewDeleteCategoryUseCase(repository *c_repository.CategoryRepository) *DeleteCategoryUseCase {
	return &DeleteCategoryUseCase{repository}
}

func (u *DeleteCategoryUseCase) Execute(id uint) error {
	err := u.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
