package repository

import (
	"context"

	"github.com/TrevorEdris/api-template/app/model/item"
)

type (
	// ItemRepoLocal stores item models in local memory.
	ItemRepoLocal struct {
		storage map[string]item.Model
	}
)

// NewItemRepoLocal creates a new ItemRepo using local memory as the storage medium.
func NewItemRepoLocal() *ItemRepoLocal {
	return &ItemRepoLocal{
		storage: make(map[string]item.Model),
	}
}

// Get retrieves the item identified by the specified id.
func (r *ItemRepoLocal) Get(ctx context.Context, id string) (item.Model, error) {
	return item.Model{}, nil
}

// Create creates a new item with the properties of the given item model.
func (r *ItemRepoLocal) Create(ctx context.Context, it item.Model) (item.Model, error) {
	return item.Model{}, nil
}

// Create updates the fields of the item identified by id to match the fields of the given item model.
func (r *ItemRepoLocal) Update(ctx context.Context, id string, updates item.Model) (item.Model, error) {
	return item.Model{}, nil
}

// Delete removes the item identified by the specified id.
func (r *ItemRepoLocal) Delete(ctx context.Context, id string) error {
	return nil
}
