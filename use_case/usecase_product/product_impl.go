package usecase_product

import (
	"errors"

	"gorm.io/gorm"
	"main.go/helper"
	"main.go/models"
	"main.go/models/dto/product_dto"
	"main.go/models/entity/product"
)

func (product_usecase *productUsecase) GetProducts() models.Response {
	productlist, err := product_usecase.productRepo.GetProducts()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}

	response := []product_dto.ProductList{}
	for _, products := range productlist {
		resProduct := product_dto.ProductList{
			ID:          products.ID,
			Name:        products.Name,
			Description: products.Description,
			Status:      products.Status,
		}
		response = append(response, resProduct)
	}

	return helper.ResponseSuccess("ok", nil, response, 200)
}

func (product_usecase *productUsecase) GetProduct(id string) models.Response {
	userData, err := product_usecase.productRepo.GetProduct(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}
	maker := models.Action{
		ID:   userData.MakerID,
		Name: userData.MakerName,
	}
	checker := models.Action{
		ID:   userData.CheckerID,
		Name: userData.CheckerName,
	}
	signer := models.Action{
		ID:   userData.SignerID,
		Name: userData.SignerName,
	}
	response := product_dto.ProductDetail{
		ID:          userData.ID,
		Name:        userData.Name,
		Description: userData.Description,
		Status:      userData.Status,
		Maker:       maker,
		Checker:     checker,
		Signer:      signer,
	}

	return helper.ResponseSuccess("ok", nil, response, 200)
}

func (products *productUsecase) CreateProduct(newProduct product_dto.Product) models.Response {
	userInsert := product.Product{
		ID:          newProduct.ID,
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Status:      "inactive",
	}

	userData, err := products.productRepo.CreateProduct(userInsert)

	if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}

	return helper.ResponseSuccess("ok", nil, map[string]interface{}{
		"id": userData.ID}, 201)
}

func (products *productUsecase) UpdateProduct(productUpdate product_dto.Product, id string, actionType string) models.Response {
	productInsert := product.Product{
		Name:        productUpdate.Name,
		Description: productUpdate.Description,
	}
	_, err := products.productRepo.UpdateProduct(productInsert, id, actionType)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}
	return helper.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (products *productUsecase) DeleteProduct(id string) models.Response {

	err := products.productRepo.DeleteProduct(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}
	return helper.ResponseSuccess("ok", nil, nil, 200)
}
