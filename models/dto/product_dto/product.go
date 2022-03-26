package product_dto

import "main.go/models"

type Product struct {
	ID          string      `gorm:"primaryKey" json:"id"`
	Name        string      `gorm:"size:255" json:"name"`
	Description string      `json:"description"`
	Status      string      `gorm:"size:255" json:"status"`
	Maker       models.Role `gorm:"foreignKey:MakerID;references:ID" json:"maker"`
	MakerID     string      `json:"maker_id"`
	Checker     models.Role `gorm:"foreignKey:CheckerID;references:ID" json:"checker"`
	CheckerID   string      `json:"checker_id"`
	Signer      models.Role `gorm:"foreignKey:SignerID;references:ID" json:"signer"`
	SignerID    string      `json:"signer_id"`
	CreatedAT   string      `json:"created_at"`
	UpdatedAT   string      `json:"updated_at"`
	DeletedAT   string      `json:"deleted_at"`
}
