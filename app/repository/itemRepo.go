package repository

import (
	"github.com/labstack/echo/v4"

	"github.com/TrevorEdris/api-template/app/domain"
	"github.com/TrevorEdris/api-template/app/infrastructure"
)

type (
	ItemRepository interface {
		GetByID(ectx echo.Context, id int) (domain.Item, error)
	}

	itemRepository struct {
		db infrastructure.PostgresDriver
	}

	dbItem struct {
		Name        string  `db:"name"`
		Description string  `db:"description"`
		ID          int     `db:"id"`
		Price       float64 `db:"price"`
	}
)

func NewItemRepository(db infrastructure.PostgresDriver) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) GetByID(ectx echo.Context, id int) (domain.Item, error) {
	result := dbItem{}
	err := r.db.GetContext(ectx.Request().Context(), &result, "SELECT * FROM items WHERE id=$1", id)
	if err != nil {
		return domain.Item{}, err
	}

	return r.convertToDomain(result), nil
}

func (r *itemRepository) convertToDomain(i dbItem) domain.Item {
	return domain.Item{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
	}
}
