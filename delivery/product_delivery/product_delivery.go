package product_delivery

import (
	"github.com/gin-gonic/gin"
	"main.go/use_case/usecase_product"
)

type ProductDelivery interface {
	GetProducts(*gin.Context)
	GetProduct(*gin.Context)
	CreateProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
}

type productDelivery struct {
	productUsecase usecase_product.ProductUsecase
}

func GetProductDelivery(usecase usecase_product.ProductUsecase) ProductDelivery {
	return &productDelivery{
		productUsecase: usecase,
	}
}
