package controllers

import (
	"net/http"

	"github.com/diyor200/ecommerce/db"
	"github.com/gin-gonic/gin"
)

func (h handler) GetProducts(ctx *gin.Context) {
	products, err := db.GetPruducts(h.DB)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(200, gin.H{"products": products})
}
