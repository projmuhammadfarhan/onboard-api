package usecase_user

import (
	"main.go/models"
	"main.go/models/dto/login_dto"
	"main.go/models/dto/user_dto"
	"main.go/repository/user_repo"
	"main.go/use_case/jwt_usecase"
)

type UserUsecase interface {
	UserLogin(login_dto.UserLogin) models.Response
	GetUsers() models.Response
	GetUser(string) models.Response
	CreateUser(user_dto.User) models.Response
	UpdateUser(user_dto.User, string) models.Response
	DeleteUser(string) models.Response
}

type userUsecase struct {
	userRepo   user_repo.UserRepository
	jwtUsecase jwt_usecase.JwtUsecase
}

func GetUserUsecase(userRepository user_repo.UserRepository, jwtAuthen jwt_usecase.JwtUsecase) UserUsecase {
	return &userUsecase{
		userRepo:   userRepository,
		jwtUsecase: jwtAuthen,
	}
}
