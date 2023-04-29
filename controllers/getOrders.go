package controllers

import (
	"fmt"
	"net/http"

	"github.com/diyor200/ecommerce/db"
	"github.com/gin-gonic/gin"
)

func (h handler) GetOrders(ctx *gin.Context) {
	orders, err := db.GetOrders(h.DB)

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	fmt.Println(orders)
	ctx.JSON(200, orders)
}
