package local

import (
	"context"

	"github.com/TrevorEdris/api-template/app/model/item"
)

type (
	// ItemRepo stores item models in local memory.
	ItemRepo struct {
		storage map[string]item.Model
	}
)

// NewItemRepo creates a new ItemRepo using local memory as the storage medium.
func NewItemRepo() *ItemRepo {
	return &ItemRepo{
		storage: make(map[string]item.Model),
	}
}

// Get retrieves the item identified by the specified id.
func (r *ItemRepo) Get(ctx context.Context, id string) (item.Model, error) {
	return item.Model{}, nil
}

// Create creates a new item with the properties of the given item model.
func (r *ItemRepo) Create(ctx context.Context, it item.Model) (item.Model, error) {
	return item.Model{}, nil
}

// Create updates the fields of the item identified by id to match the fields of the given item model.
func (r *ItemRepo) Update(ctx context.Context, id string, updates item.Model) (item.Model, error) {
	return item.Model{}, nil
}

// Delete removes the item identified by the specified id.
func (r *ItemRepo) Delete(ctx context.Context, id string) error {
	return nil
}
