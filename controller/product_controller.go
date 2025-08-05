package controller

import (
	"api-go/model"
	"api-go/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	//usecase
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			Id:    1,
			Name:  "Batata",
			Price: 25,
		},
	}
	ctx.JSON(http.StatusOK, products)
}
