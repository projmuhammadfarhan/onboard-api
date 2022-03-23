package role_repo

import (
	"gorm.io/gorm"
)

type RoleRepository interface {
	// GetRoles() (models.RoleFull, error)
	// GetRole(string) (*product.ProductDetail, error)
	// CreateProduct(product.Product) (*product.Product, error)
	// UpdateProduct(product.Product, string, string) (*product.Product, error)
	// DeleteProduct(string) error
}

type roleRepository struct {
	mysqlConnection *gorm.DB
}

func GetRoleRepository(mysqlConn *gorm.DB) RoleRepository {
	return &roleRepository{
		mysqlConnection: mysqlConn,
	}
}
