package repository

import (
	"context"

	"github.com/TrevorEdris/api-template/app/models/item"
)

type databaseClient interface {
    // TODO: In a real implementation, these would have some well-defined input formats and output formats.
    GetItem(ctx context.Context, databaseInputFormat string) error
    CreateItem(ctx context.Context, databaseInputFormat string) error
    UpdateItem(ctx context.Context, databaseInputFormat string) error
    DeleteItem(ctx context.Context, databaseInputFormat string) error
}

// ItemSomeDatabase is an implementation of a Repository that uses SomeDatabse as the backend.
type ItemSomeDatabase struct {
    // TODO: In a real implementation, this would be the actual client
    // TODO: for connecting to the database.
    dbClient databaseClient
}

// NewItemSomeDatabase returns an instance of the item repository.
func NewItemSomeDatabase(dbClient databaseClient) *ItemSomeDatabase {
    return &ItemSomeDatabase{dbClient: dbClient}
}

// Get retrieves the item specified by itemID.
func (db ItemSomeDatabase) Get(ctx context.Context, itemID string) (item.Item, error) {
    err := db.dbClient.GetItem(ctx, itemID)
    return item.Item{}, err
}

// Put puts the item into the database.
func (db ItemSomeDatabase) Put(ctx context.Context, item item.Item) error {
    err := db.dbClient.CreateItem(ctx, item.ID)
    return err
}

// Update overwrites the item in the database.
func (db ItemSomeDatabase) Update(ctx context.Context, item item.Item) error {
    err := db.dbClient.UpdateItem(ctx, item.ID)
    return err
}

// Delete deletes the item specified by itemID.
func (db ItemSomeDatabase) Delete(ctx context.Context, itemID string) error {
    err := db.dbClient.DeleteItem(ctx, itemID)
    return err
}
