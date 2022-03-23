package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/delivery/product_delivery"
	"main.go/delivery/user_delivery"
	"main.go/repository/product_repo"
	"main.go/repository/user_repo"
	"main.go/use_case/usecase_product"
	"main.go/use_case/usecase_user"
)

func HandlerRequest() {
	config.InitConfig()
	connection := config.Connect()
	productRepository := product_repo.GetProductRepository(connection)
	productUsecase := usecase_product.GetProductUsecase(productRepository)
	productDelivery := product_delivery.GetProductDelivery(productUsecase)
	userRepository := user_repo.GetUserRepository(connection)
	userUsecase := usecase_user.GetUserUsecase(userRepository)
	userDelivery := user_delivery.GetUserDelivery(userUsecase)
	// roleRepository := role_repo.GetRoleRepository(connection)
	router := gin.Default()

	router.Use(cors.Default())

	// Product Router
	router.GET("/products", productDelivery.GetProducts)
	router.GET("/products/:id", productDelivery.GetProduct)
	router.POST("/products", productDelivery.CreateProduct)
	router.PUT("/products/:id", productDelivery.UpdateProduct)
	router.PUT("/products/:id/:type/", productDelivery.UpdateProduct)
	router.DELETE("/products/:id", productDelivery.DeleteProduct)

	// User Login Router
	router.POST("/login", userDelivery.UserLogin)

	// User Router
	router.GET("/users", userDelivery.GetUsers)
	router.GET("/users/:id", userDelivery.GetUser)
	router.POST("/users", userDelivery.CreateUser)
	router.PUT("/users/:id", userDelivery.UpdateUser)
	router.DELETE("/users/:id", userDelivery.DeleteUser)

	// RUN ROUTER
	router.Run(":8001")
}
