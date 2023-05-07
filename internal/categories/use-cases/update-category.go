package use_cases

import (
	"errors"

	c_repository "github.com/GabrielBrotas/go-categories-msvc/internal/categories/repository"
)

type UpdateCategoryUseCase struct {
	repository *c_repository.CategoryRepository
}

func NewUpdateCategoryUseCase(repository *c_repository.CategoryRepository) *UpdateCategoryUseCase {
	return &UpdateCategoryUseCase{repository}
}

type UpdateCategoryInput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (u *UpdateCategoryUseCase) Execute(input UpdateCategoryInput) error {
	category, err := u.repository.FindById(input.ID)

	if err != nil {
		return err
	}

	if category == nil {
		return errors.New("category not found")
	}

	if category.Name == input.Name {
		return nil
	}

	category_with_same_name, err := u.repository.FindByName(category.Name)

	if err != nil {
		return err
	}

	if category_with_same_name != nil && category_with_same_name.ID != input.ID {
		return errors.New("name already in use")
	}

	category, err = category.UpdateName(input.Name)

	if err != nil {
		return err
	}

	err = u.repository.Update(category)

	if err != nil {
		return err
	}

	return nil
}
