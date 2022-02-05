//go:build unit
// +build unit

package dynamodb

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/TrevorEdris/api-template/app/config"
)

type itemSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	mockDDB *mockDynamodbClient
	storage *ItemRepo
}

func TestItem(t *testing.T) {
	suite.Run(t, &itemSuite{})
}

func (s *itemSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockDDB = NewmockDynamodbClient(s.ctrl)
	s.storage = NewItemRepo(&config.Config{
		DynamoDB: config.DynamoDB{
			ItemTable: "testTableName",
		},
	}, s.mockDDB)
}

func (s *itemSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *itemSuite) TestItem_GetItem_NotFound() {
	testID := "1234"
	testErr := errors.New("GetItem error")
	s.mockDDB.EXPECT().GetItem(gomock.Any(), &dynamodb.GetItemInput{
		Key:       map[string]types.AttributeValue{"id": &types.AttributeValueMemberS{Value: testID}},
		TableName: aws.String(s.storage.table),
	}).Return(nil, testErr)

	result, err := s.storage.Get(context.Background(), testID)
	assert.ErrorIs(s.T(), err, testErr)
	assert.Empty(s.T(), result)
}
