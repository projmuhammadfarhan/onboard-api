package role_repo

import (
	"gorm.io/gorm"
	"main.go/models"
)

func (repo *roleRepository) GetRoles() ([]models.Role, error) {
	roles := []models.Role{}
	err := repo.mysqlConnection.Model(&models.Role{}).Scan(&roles).Error
	if err != nil {
		return nil, err
	}

	if len(roles) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return roles, nil
}
