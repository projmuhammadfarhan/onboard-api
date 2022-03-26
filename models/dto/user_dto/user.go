package user_dto

import (
	"main.go/models"
)

type User struct {
	ID             string              `json:"id"`
	Name           string              `json:"name" binding:"required" gorm:"type=varchar(255)"`
	PersonalNumber string              `json:"personalNumber" binding:"required" gorm:"type=varchar(255)"`
	Email          string              `json:"email" binding:"required" gorm:"type=varchar(255)"`
	RoleID         models.RoleResponse `json:"role"`
	Active         bool                `json:"active"`
	Password       string              `json:"password" binding:"required" gorm:"type=varchar(255)"`
}
