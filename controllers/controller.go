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

	// GET handlers
	routes.GET("/products", h.GetProducts)
	routes.GET("/download", h.DownloadExcel)

	// POST handlers
	routes.POST("/add/product", h.AddProduct)
	routes.POST("/update/product", h.UpdateProduct)
	routes.POST("/buy", h.BuyProduct)

	// DELETE handlers
	routes.DELETE("/delete/:id", h.DeleteProduct)

	// swagger handler
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
