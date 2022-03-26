package user_dto

import (
	"main.go/models"
)

type UserList struct {
	ID     string              `json:"id"`
	Name   string              `json:"name"`
	RoleID models.RoleResponse `json:"role"`
	Active bool                `json:"active"`
}
