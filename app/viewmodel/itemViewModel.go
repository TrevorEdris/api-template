package viewmodel

import "github.com/TrevorEdris/api-template/app/domain"

type (
	ItemGetResponse struct {
		Item
	}

	Item struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		ID          int     `json:"id"`
		Price       float64 `json:"price"`
	}
)

func NewItemGetResponse(i domain.Item) ItemGetResponse {
	return ItemGetResponse{
		Item{
			ID:          i.ID,
			Name:        i.Name,
			Description: i.Description,
			Price:       i.Price,
		},
	}
}
