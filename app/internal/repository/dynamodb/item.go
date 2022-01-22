package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/TrevorEdris/api-template/app/config"
	"github.com/TrevorEdris/api-template/app/model/item"
)

type (
	ItemRepo struct {
		storage *dynamodb.Client
	}
)

func NewItemRepo(cfg *config.Config) *ItemRepo {
	return &ItemRepo{
		storage: dynamodb.NewFromConfig(cfg.AWS.AWSCfg),
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
