package controllers

import (
	"net/http"

	"github.com/diyor200/ecommerce/db"
	"github.com/diyor200/ecommerce/model"
	"github.com/gin-gonic/gin"
)

type AddProductRequestBody struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (h handler) AddProduct(ctx *gin.Context) {
	var product model.Product
	body := AddProductRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product.Name = body.Name
	product.Price = body.Price

	if result := db.AddProduct(product, h.DB); result != nil {
		ctx.AbortWithError(http.StatusBadRequest, result)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Muvaffaqiyatli qo'shildi!"})

}
