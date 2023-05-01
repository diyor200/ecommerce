package controllers

import (
	"github.com/diyor200/ecommerce/db"
	"github.com/diyor200/ecommerce/model"
	"github.com/gin-gonic/gin"
)

// @Summary Update a product
// @Description Update an existing product by ID
// @Tags Products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID to update"
// @Param product body model.Product true "Updated product details"
// @Success 201 {string} string "message: Muvaffaqiyatli o'zgartirildi!"
// @Failure 404 {string} string "message: Bunday id dagi product yo'q!"
// @Router /api/products/{id} [post]
func (h handler) UpdateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.AbortWithError(404, err)
		return
	}
	if err := db.UpdateProduct(product, h.DB); err != nil {
		ctx.JSON(404, "Bunday id dagi product yo'q!")
		return
	}

	ctx.JSON(201, gin.H{"message": "Muvaffaqiyatli o'zgartirildi!"})

}
