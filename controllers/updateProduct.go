package controllers

import (
	"github.com/diyor200/ecommerce/db"
	"github.com/diyor200/ecommerce/model"
	"github.com/gin-gonic/gin"
)

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
