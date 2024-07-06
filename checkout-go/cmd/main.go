package main

import (
	"checkout/controller"
	"checkout/repository"
	"checkout/services/queue"
	"checkout/usecase"
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

	queue := queue.Connect()

	ProductRepository := repository.NewCheckoutRepository(queue)
	ProductUseCase := usecase.NewCheckoutUsecase(ProductRepository)
	productController := controller.NewCheckoutController(ProductUseCase)

	server.POST("/checkout", productController.Checkout)

	server.Run(os.Getenv("PORT"))
}
