package controllers

import (
	"net/http"

	"github.com/diyor200/ecommerce/db"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// @Summary Get Products
// @Description Get a list of all products
// @Tags Products
// @Produce json
// @Success 200
// @Failure 404 "error"
// @Router /api/products [get]
func (h handler) GetProducts(ctx *gin.Context) {
	products, err := db.GetPruducts(h.DB)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(200, gin.H{"products": products})
}
