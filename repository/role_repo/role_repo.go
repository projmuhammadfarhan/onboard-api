package role_repo

import (
	"gorm.io/gorm"
	"main.go/models"
)

type RoleRepository interface {
	GetRoles() ([]models.Role, error)
}

type roleRepository struct {
	mysqlConnection *gorm.DB
}

func GetRoleRepository(mysqlConn *gorm.DB) RoleRepository {
	return &roleRepository{
		mysqlConnection: mysqlConn,
	}
}
