package usecase_role

import (
	"main.go/models"
	"main.go/repository/role_repo"
)

type RoleUsecase interface {
	GetRoles() models.Response
}

type roleUsecase struct {
	roleRepo role_repo.RoleRepository
}

func GetRoleUsecase(roleRepository role_repo.RoleRepository) RoleUsecase {
	return &roleUsecase{
		roleRepo: roleRepository,
	}
}
