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

/*
*

@Summary Adds a new product to the database.
@Description This endpoint receives a JSON payload containing the product name and price,
and adds the product to the database if the payload is valid.
The endpoint returns a success message if the product is added successfully.
@Tags Products
@Accept json
@Produce json
@Param request body AddProductRequestBody true "Product object that needs to be added to the database"
@Success 201 {object} gin.H{"message": "Muvaffaqiyatli qo'shildi!"}
@Failure 400 {object} error "Bad Request. The request could not be processed."
@Router /products [post]
*/
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
