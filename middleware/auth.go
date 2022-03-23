package middleware

import (
	"github.com/gin-gonic/gin"
	"main.go/helper"
	"main.go/use_case/jwt_usecase"
)

func JWTAuth(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader)
		if err != nil {
			resp := helper.ResponseError("You are unathorized", err, 401)
			c.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}

		c.Set("user_id", userId)
	}
}
