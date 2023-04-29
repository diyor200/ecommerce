package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type handler struct {
	DB *sql.DB
}

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/api")
	routes.GET("/products", h.GetProducts)
	routes.POST("/add/product", h.AddProduct)
	routes.POST("/update/product", h.UpdateProduct)
	routes.DELETE("/delete/:id", h.DeleteProduct)
	routes.GET("/orders", h.GetOrders)
	routes.POST("/buy", h.BuyProduct)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
