package usecase_product

import (
	"main.go/models"
	"main.go/models/dto/product_dto"
	"main.go/repository/product_repo"
)

type ProductUsecase interface {
	GetProducts() models.Response
	GetProduct(string) models.Response
	CreateProduct(product_dto.Product) models.Response
	UpdateProduct(product_dto.Product, string, string) models.Response
	DeleteProduct(string) models.Response
}

type productUsecase struct {
	productRepo product_repo.ProductRepository
}

func GetProductUsecase(productRepository product_repo.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: productRepository,
	}
}
