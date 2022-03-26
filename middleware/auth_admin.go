package middleware

// func JWTAuthAdmin(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		userId, role, err := jwtUsecase.ValidateTokenAndGetRole(authHeader)
// 		if err != nil {
// 			resp := helper.ResponseError("You are unathorized", "Invalid Token", 401)
// 			c.AbortWithStatusJSON(resp.StatusCode, resp)
// 			return
// 		}

// 		if role != "admin" {
// 			resp := helper.ResponseError("Wrong Access", "You can't Access this Action", 401)
// 			c.AbortWithStatusJSON(resp.StatusCode, resp)
// 			return
// 		}

// 		c.Set("user_id", userId)
// 	}
// }
