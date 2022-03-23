package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID             string `gorm:"primaryKey"`
	PersonalNumber string `json:"personal_number"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	RoleID         string `json:"role_id"`
	Active         bool   `json:"active"`
}
