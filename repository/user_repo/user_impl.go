package user_repo

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"main.go/helper"
	"main.go/models"
	"main.go/models/dto/user_dto"
	"main.go/models/entity/user"
)

func (repo *userRepository) GetRoleByRoleId(id string) (*models.Role, error) {
	role := models.Role{}
	result := repo.mysqlConnection.Where("id = ?", id).Find(&role)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &role, nil
}

func (repo *userRepository) GetUserByPN(pn string) (*user.User, error) {
	user := user.User{}
	result := repo.mysqlConnection.Where("personal_number = ?", pn).Find(&user)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func (repo *userRepository) GetUsers() ([]user.UserList, error) {
	users := []user.UserList{}
	err := repo.mysqlConnection.Model(&user.User{}).Select("users.name, users.active, users.id, roles.title, users.role_id").Joins("left join roles on roles.id = users.role_id").Scan(&users).Error
	// check := repo.mysqlConnection.Unscoped().Where("id = ?", "d9e9e7bc-4f9a-4572-a436-d1a764d73d45").Find(&user.User{})
	// fmt.Println("data deleted :", check)
	if err != nil {
		return nil, err
	}

	if len(users) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return users, nil
}

func (repo *userRepository) GetUser(id string) (*user_dto.UserDetail, error) {
	users := user_dto.UserDetail{}
	err := repo.mysqlConnection.Model(&user.User{}).Where("users.id = ?", id).Select("users.name, users.active, users.email, users.personal_number, users.id, roles.title, users.role_id").Joins("left join roles on roles.id = users.role_id").Scan(&users).Error
	if err != nil {
		return nil, err
	}

	if (users == user_dto.UserDetail{}) {
		return nil, gorm.ErrRecordNotFound
	}

	return &users, nil
}

func (repo *userRepository) CreateUser(user user.User) (*user.User, *models.Role, error) {
	role := models.Role{}

	result := repo.mysqlConnection.Where("personal_number = ?", user.PersonalNumber).Find(&user)
	if result.RowsAffected > 0 {
		return nil, nil, gorm.ErrRegistered
	}

	user.ID = uuid.New().String()
	hash, _ := helper.HashPassword(user.Password)
	user.Password = hash

	if err := repo.mysqlConnection.Where("title = ?", "admin").Find(&role).Error; err != nil {
		return nil, nil, err
	}

	user.RoleID = role.ID

	if err := repo.mysqlConnection.Create(&user).Error; err != nil {
		return nil, nil, err
	}

	return &user, &role, nil
}

func (repo *userRepository) UpdateUser(user user.User, id string) (*user.User, error) {
	result := repo.mysqlConnection.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"name":            user.Name,
		"password":        user.Password,
		"role_id":         user.RoleID,
		"active":          user.Active,
		"email":           user.Email,
		"personal_number": user.PersonalNumber,
	})

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func (repo *userRepository) DeleteUser(id string) error {
	// get := repo.mysqlConnection
	// if err := get.Where("id = ?", id).Delete(user.User{}).Error; err != nil {
	// sql = fmt.Sprintf("%s WHERE id = '%s'", sql, id)
	fmt.Println("CHECK ID :", id)
	if err := repo.mysqlConnection.Delete(&user.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	// if err := repo.mysqlConnection.Delete(&entity.User{}, id).Error; err != nil  {
	// 	return err
	// }
	return nil
}
