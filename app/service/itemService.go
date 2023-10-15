package service

import (
	"github.com/TrevorEdris/api-template/app/domain"
	"github.com/TrevorEdris/api-template/app/repository"
	"github.com/labstack/echo/v4"
)

type (
	ItemService interface {
		GetByID(ectx echo.Context, id int) (domain.Item, error)
	}

	itemService struct {
		repo repository.ItemRepository
	}
)

func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{
		repo: repo,
	}
}

func (s *itemService) GetByID(ectx echo.Context, id int) (domain.Item, error) {
	return s.repo.GetByID(ectx, id)
}
