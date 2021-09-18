package somedatabase

import (
	"context"
)

// TODO: In a real implementation, this would be something like dynamoDB or mongoDB
// TODO: and would be imported from that database's SDK.
type SomeDatabaseClient struct {}

func (c *SomeDatabaseClient) GetItem(ctx context.Context, databaseInputFormat string) error {return nil}
func (c *SomeDatabaseClient) CreateItem(ctx context.Context, databaseInputFormat string) error {return nil}
func (c *SomeDatabaseClient) UpdateItem(ctx context.Context, databaseInputFormat string) error {return nil}
func (c *SomeDatabaseClient) DeleteItem(ctx context.Context, databaseInputFormat string) error {return nil}

func NewSomeDatabaseClient(cfg string) SomeDatabaseClient {
    return SomeDatabaseClient{}
}
