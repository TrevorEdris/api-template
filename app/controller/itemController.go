package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TrevorEdris/api-template/app/service"
	"github.com/TrevorEdris/api-template/app/viewmodel"
	"github.com/labstack/echo/v4"
)

type (
	ItemController interface {
		GetOne(ectx echo.Context) error
	}

	itemController struct {
		svc service.ItemService
	}
)

func NewItemController(svc service.ItemService) ItemController {
	return &itemController{
		svc: svc,
	}
}

func (c *itemController) GetOne(ectx echo.Context) error {
	idParam := ectx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return fmt.Errorf("id supplied '%s' was not an int: %w", idParam, err)
	}

	item, err := c.svc.GetByID(ectx, id)
	if err != nil {
		return err
	}
	ectx.Logger().Debugf("Retrieved item %v", item)

	return ectx.JSON(http.StatusOK, viewmodel.NewItemGetResponse(item))
}
