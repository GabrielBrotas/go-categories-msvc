package models_test

import (
	"testing"

	categories "github.com/GabrielBrotas/go-categories-msvc/internal/categories/models"
)

func TestNewCategory(t *testing.T) {
	category, err := categories.NewCategory("Test")

	if err == nil {
		t.Errorf("category.IsValid() not working, name must be greater than 5. Got %d", len(category.Name))
	}

	category, err = categories.NewCategory("Test2")

	if err != nil {
		t.Errorf("category.IsValid() not working, name is be greater than 5. Got %d", len(category.Name))
	}
}
