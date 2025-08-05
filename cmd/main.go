package main

import (
	"api-go/controller"
	"api-go/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	productUseCase := usecase.NewProductUseCase()
	productController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.Run(":8080")
}
