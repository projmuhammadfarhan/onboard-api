package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/delivery/product_delivery"
	"main.go/delivery/role_delivery"
	"main.go/delivery/user_delivery"
	"main.go/middleware"
	"main.go/repository/product_repo"
	"main.go/repository/role_repo"
	"main.go/repository/user_repo"
	"main.go/use_case/jwt_usecase"
	"main.go/use_case/usecase_product"
	"main.go/use_case/usecase_role"
	"main.go/use_case/usecase_user"
)

func HandlerRequest() {
	config.InitConfig()
	connection := config.Connect()

	// User
	userRepository := user_repo.GetUserRepository(connection)
	jwtUsecase := jwt_usecase.GetJwtUsecase(userRepository)
	userUsecase := usecase_user.GetUserUsecase(userRepository, jwtUsecase)
	userDelivery := user_delivery.GetUserDelivery(userUsecase)

	// Role
	roleRepository := role_repo.GetRoleRepository(connection)
	roleUsecase := usecase_role.GetRoleUsecase(roleRepository)
	roleDelivery := role_delivery.GetRoleDelivery(roleUsecase)

	// Product
	productRepository := product_repo.GetProductRepository(connection)
	productUsecase := usecase_product.GetProductUsecase(productRepository)
	productDelivery := product_delivery.GetProductDelivery(productUsecase)

	router := gin.Default()
	router.Use(cors.Default())

	// Users
	router.POST("/users", middleware.JWTAuthAdmin(jwtUsecase, []string{"admin", "maker"}), userDelivery.CreateUser)
	router.POST("/new/users", userDelivery.CreateUser)
	router.PUT("/users/:id", middleware.JWTAuthAdmin(jwtUsecase, []string{"admin", "checker", "signer"}), userDelivery.UpdateUser)
	router.DELETE("/users/:id", middleware.JWTAuthAdmin(jwtUsecase, []string{"admin"}), userDelivery.DeleteUser)
	router.GET("/users", middleware.JWTAuthAdmin(jwtUsecase, []string{"admin", "viewer", "checker", "signer", "maker"}), userDelivery.GetUsers)
	router.GET("/users/:id", middleware.JWTAuthAdmin(jwtUsecase, []string{"admin", "viewer", "checker", "signer", "maker"}), userDelivery.GetUser)

	// Product
	router.POST("/products", middleware.JWTAuthAdmin(jwtUsecase, []string{"admin", "viewer", "checker", "signer", "maker"}), productDelivery.CreateProduct)
	router.PUT("/products/:id", middleware.JWTAuthAdmin(jwtUsecase, []string{"admin", "viewer", "checker", "signer", "maker"}), productDelivery.UpdateProduct)
	router.DELETE("/products/:id", middleware.JWTAuthAdmin(jwtUsecase, []string{"admin"}), productDelivery.DeleteProduct)
	router.GET("/products", middleware.JWTAuthAdmin(jwtUsecase, []string{"admin", "viewer", "checker", "signer", "maker"}), productDelivery.GetProducts)
	router.GET("/products/:id", middleware.JWTAuthAdmin(jwtUsecase, []string{"admin", "viewer", "checker", "signer", "maker"}), productDelivery.GetProduct)
	//Product Checked
	router.GET("/products/:id/checked", middleware.JWTAuthAdmin(jwtUsecase, []string{"admin", "checker"}), productDelivery.GetProduct)
	//Product Published
	router.GET("/products/:id/published", middleware.JWTAuthAdmin(jwtUsecase, []string{"admin", "signer"}), productDelivery.GetProduct)

	// Role Router
	router.GET("/roles", roleDelivery.GetRoles)

	// Login
	router.POST("/login", userDelivery.UserLogin)

	// RUN ROUTER
	router.Run(":8001")
}
