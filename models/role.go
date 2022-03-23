package models

type Role struct {
	ID    string `gorm:"primaryKey"`
	Title string `json:"title"`
}
