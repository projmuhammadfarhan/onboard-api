package role_delivery

import (
	"github.com/gin-gonic/gin"
	"main.go/use_case/usecase_role"
)

type RoleDelivery interface {
	GetRoles(*gin.Context)
}

type roleDelivery struct {
	usecase usecase_role.RoleUsecase
}

func GetRoleDelivery(roleUsecase usecase_role.RoleUsecase) RoleDelivery {
	return &roleDelivery{
		usecase: roleUsecase,
	}
}
