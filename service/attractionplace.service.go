package service

import (
	"attractionplace/message"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateAttractionPlaceService interface {
	CreatePlace(c echo.Context) error
}

type createAttractionPlaceService struct{}

func NewCreateAttractionPlaceService() CreateAttractionPlaceService {
	return createAttractionPlaceService{}
}

func (caps createAttractionPlaceService) CreatePlace(c echo.Context) error {
	p := message.AttractionPlace{}
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, p)
}
