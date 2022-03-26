package product

import (
	"gorm.io/gorm"
	"main.go/models/entity/user"
)

type Product struct {
	gorm.Model
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255" json:"name"`
	Description string    `json:"description"`
	Status      string    `gorm:"size:255" json:"status"`
	Maker       user.User `gorm:"foreignKey:MakerID;association_foreignKey:ID"`
	MakerID     string    `gorm:"size:191" json:"maker_id"`
	Checker     user.User `gorm:"foreignKey:CheckerID;association_foreignKey:ID"`
	CheckerID   string    `gorm:"size:191" json:"checker_id"`
	Signer      user.User `gorm:"foreignKey:SignerID;association_foreignKey:ID"`
	SignerID    string    `gorm:"size:191" json:"signer_id"`
}
