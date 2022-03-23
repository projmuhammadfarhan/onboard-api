package user_repo

import (
	"gorm.io/gorm"
	"main.go/models"
	"main.go/models/entity/user"
)

type UserRepository interface {
	GetUserByPN(string) (*user.User, error)
	GetRoleByRoleId(string) (*models.Role, error)
	GetUsers() ([]user.UserList, error)
	GetUser(string) (*user.UserDetail, error)
	CreateUser(user.User) (*user.User, *models.Role, error)
	UpdateUser(user.User, string) (*user.User, error)
	DeleteUser(string) error
}

type userRepository struct {
	mysqlConnection *gorm.DB
}

func GetUserRepository(mysqlConn *gorm.DB) UserRepository {
	return &userRepository{
		mysqlConnection: mysqlConn,
	}
}
