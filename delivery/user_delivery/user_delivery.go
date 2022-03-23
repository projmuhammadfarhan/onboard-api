package user_delivery

import (
	"github.com/gin-gonic/gin"
	"main.go/use_case/usecase_user"
)

type UserDelivery interface {
	UserLogin(*gin.Context)
	GetUsers(*gin.Context)
	GetUser(*gin.Context)
	CreateUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

type userDelivery struct {
	usecase usecase_user.UserUsecase
}

func GetUserDelivery(userUsecase usecase_user.UserUsecase) UserDelivery {
	return &userDelivery{
		usecase: userUsecase,
	}
}
