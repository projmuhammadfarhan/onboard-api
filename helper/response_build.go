package helper

import (
	"main.go/models"
)

func ResponseError(status string, err interface{}, code int) models.Response {
	return models.Response{
		StatusCode: code,
		Status:     status,
		Error:      err,
		Data:       nil,
	}
}

func ResponseSuccess(status string, err interface{}, data interface{}, code int) models.Response {
	return models.Response{
		StatusCode: code,
		Status:     status,
		Error:      err,
		Data:       data,
	}
}
