package repositories

import "github.com/GabrielBrotas/go-categories-msvc/internal/entities"

type inMemoryCategoryRepository struct {
	db []*entities.Category
}

func NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (r *inMemoryCategoryRepository) Save(category *entities.Category) error {
	// add an ID before saving
	r.db = append(r.db, category)

	return nil
}

func (r *inMemoryCategoryRepository) List() ([]*entities.Category, error) {
	return r.db, nil
}