package middleware

import (
	"github.com/gin-gonic/gin"
	"main.go/helper"
	"main.go/use_case/jwt_usecase"
)

// func JWTAuth(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		// fmt.Println("AUTH HEADER :", authHeader)
// 		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader)
// 		if err != nil {
// 			resp := helper.ResponseError("You are unathorized", err, 401)
// 			c.AbortWithStatusJSON(resp.StatusCode, resp)
// 			return
// 		}

// 		c.Set("user_id", userId)
// 	}
// }

func JWTAuthAdmin(jwtUsecase jwt_usecase.JwtUsecase, roleTitle []string) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		userId, role, err := jwtUsecase.ValidateTokenAndGetRole(authHeader)
		if err != nil {
			resp := helper.ResponseError("You are unathorized", "Invalid Token", 401)
			c.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}
		indikator := 0
		for _, v := range roleTitle {
			// fmt.Println("CHECK V : ", v)
			// fmt.Println("CHECK ROLE : ", role)
			if v == role {
				// resp := helper.ResponseError("Wrong Access", "You can't Access this Action", 401)
				// c.AbortWithStatusJSON(resp.StatusCode, resp)
				// return
				indikator = 1
			}
		}
		if indikator == 0 {
			resp := helper.ResponseError("unauthorized", "You can't Access this Action", 401)
			c.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}
		c.Set("user_id", userId)

		// if roleTitle != role {
		// 	resp := helper.ResponseError("Wrong Access", "You can't Access this Action", 401)
		// 	c.AbortWithStatusJSON(resp.StatusCode, resp)
		// 	return
		// }

		// if role != "admin" {
		// 	resp := helper.ResponseError("Wrong Access", "You can't Access this Action", 401)
		// 	c.AbortWithStatusJSON(resp.StatusCode, resp)
		// 	return
		// }
		// if role != "checker" {
		// 	resp := helper.ResponseError("Wrong Access", "You can't Access this Action", 401)
		// 	c.AbortWithStatusJSON(resp.StatusCode, resp)
		// 	return
		// }
		// if role != "maker" {
		// 	resp := helper.ResponseError("Wrong Access", "You can't Access this Action", 401)
		// 	c.AbortWithStatusJSON(resp.StatusCode, resp)
		// 	return
		// }
		// if role != "viewer" {
		// 	resp := helper.ResponseError("Wrong Access", "You can't Access this Action", 401)
		// 	c.AbortWithStatusJSON(resp.StatusCode, resp)
		// 	return
		// }
		// if role != "signer" {
		// 	resp := helper.ResponseError("Wrong Access", "You can't Access this Action", 401)
		// 	c.AbortWithStatusJSON(resp.StatusCode, resp)
		// 	return
		// }

		// c.Set("user_id", userId)
	}
}

// func JWTAuthMaker(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		userId, role, err := jwtUsecase.ValidateTokenAndGetRole(authHeader)
// 		if err != nil {
// 			resp := helper.ResponseError("You are unathorized", "Invalid Token", 401)
// 			c.AbortWithStatusJSON(resp.StatusCode, resp)
// 			return
// 		}

// 		if role != "maker" {
// 			resp := helper.ResponseError("Wrong Access", "You can't Access this Action", 401)
// 			c.AbortWithStatusJSON(resp.StatusCode, resp)
// 			return
// 		}

// 		c.Set("user_id", userId)
// 	}
// }

// func JWTAuthChecker(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		userId, role, err := jwtUsecase.ValidateTokenAndGetRole(authHeader)
// 		if err != nil {
// 			resp := helper.ResponseError("You are unathorized", "Invalid Token", 401)
// 			c.AbortWithStatusJSON(resp.StatusCode, resp)
// 			return
// 		}

// 		if role != "checker" {
// 			resp := helper.ResponseError("Wrong Access", "You can't Access this Action", 401)
// 			c.AbortWithStatusJSON(resp.StatusCode, resp)
// 			return
// 		}

// 		c.Set("user_id", userId)
// 	}
// }

// func JWTAuthSigner(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		userId, role, err := jwtUsecase.ValidateTokenAndGetRole(authHeader)
// 		if err != nil {
// 			resp := helper.ResponseError("You are unathorized", "Invalid Token", 401)
// 			c.AbortWithStatusJSON(resp.StatusCode, resp)
// 			return
// 		}

// 		if role != "signer" {
// 			resp := helper.ResponseError("Wrong Access", "You can't Access this Action", 401)
// 			c.AbortWithStatusJSON(resp.StatusCode, resp)
// 			return
// 		}

// 		c.Set("user_id", userId)
// 	}
// }

// func JWTAuthViewer(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		userId, role, err := jwtUsecase.ValidateTokenAndGetRole(authHeader)
// 		if err != nil {
// 			resp := helper.ResponseError("You are unathorized", "Invalid Token", 401)
// 			c.AbortWithStatusJSON(resp.StatusCode, resp)
// 			return
// 		}

// 		if role != "viewer" {
// 			resp := helper.ResponseError("Wrong Access", "You can't Access this Action", 401)
// 			c.AbortWithStatusJSON(resp.StatusCode, resp)
// 			return
// 		}

// 		c.Set("user_id", userId)
// 	}
// }
