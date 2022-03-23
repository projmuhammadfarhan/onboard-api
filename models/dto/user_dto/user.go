package user_dto

import (
	"main.go/models"
)

type User struct {
	ID             string      `json:"id"`
	Name           string      `json:"name" binding:"required" gorm:"type=varchar(255)"`
	PersonalNumber string      `json:"personal_number" binding:"required" gorm:"type=varchar(255)"`
	Email          string      `json:"email" binding:"required" gorm:"type=varchar(255)"`
	RoleID         models.Role `json:"role"`
	Active         bool        `json:"active"`
	Password       string      `json:"password" binding:"required" gorm:"type=varchar(255)"`
}
