package product_repo

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"main.go/models/entity/product"
	"main.go/models/entity/user"
)

func (repo *productRepository) GetProducts() ([]product.Product, error) {
	products := []product.Product{}
	err := repo.mysqlConnection.Model(&product.Product{}).Scan(&products).Error
	if err != nil {
		return nil, err
	}

	if len(products) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return products, nil
}

func (repo *productRepository) GetProduct(id string) (*product.ProductDetail, error) {
	productDetail := product.ProductDetail{}
	products := product.Product{}
	users := []user.User{}

	if err := repo.mysqlConnection.Where("id = ?", id).Find(&products).Error; err != nil {
		return nil, err
	}

	if (product.Product{}) == products {
		return nil, gorm.ErrRecordNotFound
	}

	productDetail.ID = products.ID
	productDetail.Name = products.Name
	productDetail.Description = products.Description
	productDetail.Status = products.Status

	err := repo.mysqlConnection.Where("id IN ?", []string{products.MakerID, products.CheckerID, products.SignerID}).Find(&users).Error
	if err != nil {
		return nil, err
	}

	fmt.Println(users)

	for _, user := range users {
		if user.ID == productDetail.MakerID {
			productDetail.MakerName = user.Name
		} else if user.ID == productDetail.CheckerID {
			productDetail.CheckerName = user.Name
		} else if user.ID == productDetail.SignerID {
			productDetail.SignerName = user.Name
		}
	}

	return &productDetail, nil
}

func (repo *productRepository) CreateProduct(product product.Product) (*product.Product, error) {
	product.ID = uuid.New().String()
	product.MakerID = "system"
	product.CheckerID = ""
	product.SignerID = ""

	if err := repo.mysqlConnection.Create(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// func (repo *productRepository) CreateProduct(product product.Product) (*product.Product, error) {
// 	users := user.User{}
// 	product.ID = uuid.New().String()

// 	result := repo.mysqlConnection.Where("id = ?", product.ID).Find(&product)
// 	if result.RowsAffected > 0 {
// 		return nil, gorm.ErrRegistered
// 	}

// 	product.MakerID = users.ID
// 	product.SignerID = users.ID
// 	product.CheckerID = users.ID

// 	// product.MakerID = "system"
// 	// product.CheckerID = ""
// 	// product.SignerID = ""

// 	if err := repo.mysqlConnection.Create(&product).Error; err != nil {
// 		return nil, err
// 	}
// 	return &product, nil
// }

func (repo *productRepository) UpdateProduct(product product.Product, id string, actionType string) (*product.Product, error) {
	switch actionType {
	case "published":
		result := repo.mysqlConnection.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"name": product.Name, "description": product.Description, "signer_id": "system", "status": "active"})
		if result.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
	case "checked":
		result := repo.mysqlConnection.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"name": product.Name, "description": product.Description, "checker_id": "system"})
		if result.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
	default:
		result := repo.mysqlConnection.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"name": product.Name, "description": product.Description})
		if result.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
	}

	return &product, nil
}

func (repo *productRepository) DeleteProduct(id string) error {
	// sql := "DELETE FROM products"
	// sql = fmt.Sprintf("%s WHERE id = '%s'", sql, id)
	// if err := repo.mysqlConnection.Raw(sql).Scan(product.Product{}).Error; err != nil {
	if err := repo.mysqlConnection.Delete(&product.Product{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
