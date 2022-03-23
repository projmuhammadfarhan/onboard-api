package models

type RoleFull struct {
	ID     string `gorm:"primaryKey"`
	Title  string `json:"title"`
	Active bool   `json:"active"`
}
