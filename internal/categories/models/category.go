package models

import (
	"fmt"
	"time"

	error_pkg "github.com/GabrielBrotas/go-categories-msvc/pkg/error"
)

type Category struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := category.IsValid()

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) IsValid() error {
	ec := error_pkg.NewErrorCollection()

	if len(c.Name) < 5 {
		ec.Add(fmt.Errorf("name must be greater than 5. Got %d", len(c.Name)))
	}

	if ec.HasErrors() {
		return ec.Throw()
	}

	return nil
}

func (c *Category) UpdateName(name string) (*Category, error) {
	ec := error_pkg.NewErrorCollection()

	if len(name) < 5 {
		ec.Add(fmt.Errorf("name must be greater than 5. Got %d", len(c.Name)))
	}

	if ec.HasErrors() {
		return nil, ec.Throw()
	}

	c.Name = name

	return c, nil
}
