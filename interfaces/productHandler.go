package interfaces

import (
	"net/http"

	"github.com/danisbagus/simple-go-ddd/application"
	"github.com/danisbagus/simple-go-ddd/domain/entity"
	"github.com/labstack/echo"
)

type productHandler struct {
	productApp application.ProductAppInterface
}

func NewProductHandler(productApp application.ProductAppInterface) productHandler {
	return productHandler{productApp}
}

func (h productHandler) Insert(c echo.Context) error {

	product := new(entity.Product)
	if err := c.Bind(product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := product.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.productApp.Insert(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{"message": "Successfully insert data"}
	return c.JSON(http.StatusOK, response)
}

func (h productHandler) List(c echo.Context) error {
	products, err := h.productApp.List()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	prouductResponse := make([]entity.ProductResponse, 0)
	for _, product := range products {
		var res entity.ProductResponse
		res.ID = product.ID.Hex()
		res.Name = product.Name
		res.CategoryIDs = product.CategoryIDs
		res.Price = product.Price

		prouductResponse = append(prouductResponse, res)
	}

	response := map[string]interface{}{"message": "Successfully get data", "data": prouductResponse}
	return c.JSON(http.StatusOK, response)
}

func (h productHandler) View(c echo.Context) error {
	productID := c.Param("id")
	product, err := h.productApp.View(productID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	productResponse := new(entity.ProductResponse)
	productResponse.ID = product.ID.Hex()
	productResponse.Name = product.Name
	productResponse.CategoryIDs = product.CategoryIDs
	productResponse.Price = product.Price

	response := map[string]interface{}{"message": "Successfully get data", "data": productResponse}
	return c.JSON(http.StatusOK, response)
}

func (h productHandler) Update(c echo.Context) error {
	productID := c.Param("id")
	product := new(entity.Product)
	if err := c.Bind(product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := product.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.productApp.Update(productID, product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{"message": "Successfully update data"}
	return c.JSON(http.StatusOK, response)
}

func (h productHandler) Delete(c echo.Context) error {
	productID := c.Param("id")
	err := h.productApp.Delete(productID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{"message": "Successfully delete data"}
	return c.JSON(http.StatusOK, response)
}
