package usecase_role

import (
	"errors"

	"gorm.io/gorm"
	"main.go/helper"
	"main.go/models"
)

func (role *roleUsecase) GetRoles() models.Response {
	rolelist, err := role.roleRepo.GetRoles()
	response := []models.RoleResponse{}
	for _, role := range rolelist {
		// role := models.RoleResponse{ID: role.ID, Title: role.Title}
		responseData := models.RoleResponse{
			ID:    role.ID,
			Title: role.Title,
		}
		response = append(response, responseData)
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}
	return helper.ResponseSuccess("ok", nil, response, 200)
}
