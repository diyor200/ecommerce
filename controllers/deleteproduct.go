package controllers

import (
	"net/http"
	"strconv"

	"github.com/diyor200/ecommerce/db"
	"github.com/gin-gonic/gin"
)

type DeleteRequestBody struct {
	ID int
}

// @Summary Delete a Product
// @Description Delete a product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200
// @Failure 404 "error"
// @Router /api/delete/{id} [delete]
func (h handler) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "id int tipida bo'lishi kerak!"})
	}

	if result := db.DeleteProduct(id, h.DB); result != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Bunday id dagi product mavjud emas!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Muvaffaqiyatli o'chirildi!"})
}
