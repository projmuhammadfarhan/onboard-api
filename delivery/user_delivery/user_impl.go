package user_delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/helper"
	"main.go/models/dto/login_dto"
	"main.go/models/dto/user_dto"
)

func (res *userDelivery) UserLogin(c *gin.Context) {
	request := login_dto.UserLogin{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helper.ResponseError("Bad Request", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}
	response := res.usecase.UserLogin(request)

	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(response.StatusCode, response)
}

func (res *userDelivery) GetUsers(c *gin.Context) {
	response := res.usecase.GetUsers()
	// fmt.Printf("%+v", response)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) GetUser(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.GetUser(id)
	if response.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusOK, response)
		return
	}

	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) CreateUser(c *gin.Context) {
	request := user_dto.User{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helper.ResponseError("Bad Request", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}
	response := res.usecase.CreateUser(request)

	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	request := user_dto.User{}
	// fmt.Println("Check :", c.ShouldBindJSON(&request))
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helper.ResponseError("Bad Request Delivery", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.UpdateUser(request, id)

	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (res *userDelivery) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.DeleteUser(id)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}
