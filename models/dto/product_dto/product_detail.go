package product_dto

import (
	"main.go/models"
)

type ProductDetail struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Status      string        `json:"status"`
	Maker       models.Action `json:"maker"`
	Checker     models.Action `json:"checker"`
	Signer      models.Action `json:"signer"`
}
