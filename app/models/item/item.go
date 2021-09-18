package item

import (
	"context"
	"fmt"
)

const (
	minIdLen        = 1
	minSomeFieldLen = 1
)

// Repository defines the interface necessary for interacting with a database.
type Repository interface {
	Get(ctx context.Context, itemID string) (Item, error)
	Put(ctx context.Context, item Item) error
	Update(ctx context.Context, item Item) error
	Delete(ctx context.Context, itemID string) error
}

// Item is the actual definition of an item.
type Item struct {
	ID        string `somedatabase:"id"`
	SomeField string `somedatabase:"some_field"`
}

// Items provides a way to interact with the model and the storage.
// The business logic of dealing with an Item should be handled here.
type Items struct {
	storage Repository
}

// New creates a new Items struct.
func New(repo Repository) *Items {
	return &Items{storage: repo}
}

// NewItem creates a new Item.
func (items *Items) NewItem(id string, someField string) (Item, error) {
	if len(id) < minIdLen {
		return Item{}, fmt.Errorf("ID (%d) must be longer than %d", len(id), minIdLen)
	}

	if len(someField) < minSomeFieldLen {
		return Item{}, fmt.Errorf("SomeField (%d) must be longer than %d", len(someField), minSomeFieldLen)
	}

	return Item{
		ID:        id,
		SomeField: someField,
	}, nil
}

// GetItem gets the item from the storage.
func (items *Items) GetItem(ctx context.Context, itemID string) (Item, error) {
	return items.storage.Get(ctx, itemID)
}

// PutItem puts the item into storage.
func (items *Items) PutItem(ctx context.Context, item Item) error {
	return items.storage.Put(ctx, item)
}

// UpdateItem updates the item in storage.
func (items *Items) UpdateItem(ctx context.Context, item Item) error {
	return items.storage.Update(ctx, item)
}

// DeleteItem deletes the item from storage.
func (items *Items) Deleteitem(ctx context.Context, itemID string) error {
	return items.storage.Delete(ctx, itemID)
}
