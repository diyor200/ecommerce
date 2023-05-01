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

// @Summary Add a new product
// @Description This API endpoint is used to add a new product to the database.
// @Tags Products
// @Accept json
// @Produce json
// @Param body body AddProductRequestBody true "Product Name and Price"
// @Success 201 {string} string "message: Muvaffaqiyatli qo'shildi!"
// @Failure 400 {string} string "message: Bad Request"
// @Router /api/add/product [post]
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
