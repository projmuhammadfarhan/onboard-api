package product_delivery

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/helper"
	"main.go/models/dto/product_dto"
)

func (product_delivery *productDelivery) CreateProduct(c *gin.Context) {
	request := product_dto.Product{}
	if err := c.ShouldBindJSON(&request); err != nil {

		errorRes := helper.ResponseError("Invalid Input", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return

	}
	response := product_delivery.productUsecase.CreateProduct(request)

	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (product_delivery *productDelivery) GetProducts(c *gin.Context) {
	response := product_delivery.productUsecase.GetProducts()
	// Cek output response
	fmt.Printf("%v", response)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(response.StatusCode, response)
}

func (product_delivery *productDelivery) GetProduct(c *gin.Context) {
	id := c.Param("id")
	response := product_delivery.productUsecase.GetProduct(id)

	if response.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusOK, response)
		return
	}
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (product_delivery *productDelivery) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	actionType := c.Param("type")

	request := product_dto.Product{}

	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helper.ResponseError("Invalid Input", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := product_delivery.productUsecase.UpdateProduct(request, id, actionType)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (product_delivery *productDelivery) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	response := product_delivery.productUsecase.DeleteProduct(id)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(response.StatusCode, response)
}
