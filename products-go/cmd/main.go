package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	server := gin.Default()

	connect, err := db.Connect()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(connect)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	productController := controller.NewProductController(ProductUseCase)

	server.GET("/products", productController.GetProduct)
	server.POST("/products", productController.CreateProduct)
	server.GET("/products/:productId", productController.GetProductByID)
	server.Run(os.Getenv("PORT"))
}
