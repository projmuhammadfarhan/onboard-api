package product_repo

import (
	"gorm.io/gorm"
	"main.go/models/entity/product"
)

type ProductRepository interface {
	GetProducts() ([]product.Product, error)
	GetProduct(string) (*product.ProductDetail, error)
	CreateProduct(product.Product) (*product.Product, error)
	UpdateProduct(product.Product, string, string) (*product.Product, error)
	DeleteProduct(string) error
}

type productRepository struct {
	mysqlConnection *gorm.DB
}

func GetProductRepository(mysqlConn *gorm.DB) ProductRepository {
	return &productRepository{
		mysqlConnection: mysqlConn,
	}
}
