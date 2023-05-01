package controllers

import (
	"fmt"
	"net/http"

	"github.com/diyor200/ecommerce/db"
	"github.com/diyor200/ecommerce/model"
	"github.com/gin-gonic/gin"
)

type BuyRequestBody struct {
	Product_id int    `json:"id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
}

// @Summary Buy a product
// @Description This API endpoint is used to buy a product.
// @Tags Products
// @Accept json
// @Produce json
// @Param id body BuyRequestBody true "Product ID and user's firstname and lastname"
// @Success 201 {string} string "message: Muvaffaqiyatli harid qilindi!"
// @Failure 400 {string} string "message: Harid qilishda muammo yuzga keldi, qaytadan urinib ko'ring !"
// @Failure 404 {string} string "message: Bunday id dagi product mavjud emas!"
// @Failure 404 {string} string "message: Not Found"
// @Router /api/buy [post]
func (h handler) BuyProduct(ctx *gin.Context) {
	body := BuyRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(404, err)
		return
	}
	fmt.Println(body)

	product, err := db.GetPruduct(body.Product_id, h.DB)
	fmt.Println(product)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	if product.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Bunday id dagi product mavjud emas!"})
		return
	}

	order := model.Orders{}
	order.User.Firstname = body.Firstname
	order.User.Lastname = body.Lastname
	order.Productname = product.Name
	order.Payment = product.Price

	if err := db.DeleteProduct(body.Product_id, h.DB); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Harid qilishda muammo yuzga keldi, qaytadan urinib ko'ring !"})
		return
	}

	if result := db.AddOrders(order, h.DB); result != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": result})
		return
	}

	ctx.JSON(201, gin.H{"message": "Muvaffaqiyatli harid qilindi!"})
}
