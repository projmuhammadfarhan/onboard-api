package user

import (
	"gorm.io/gorm"
	"main.go/models"
)

type User struct {
	gorm.Model
	ID             string      `gorm:"primaryKey"`
	PersonalNumber string      `gorm:"size:255" json:"personalNumber"`
	Name           string      `gorm:"size:255" json:"name"`
	Password       string      `gorm:"size:255" json:"password"`
	Email          string      `gorm:"size:255" json:"email"`
	Role           models.Role `gorm:"foreignKey:RoleID;references:ID" `
	RoleID         string      `gorm:"size:191" json:"roleId"`
	Active         bool        `json:"active"`
}
