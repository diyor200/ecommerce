package main

import (
	"database/sql"
	"log"

	"github.com/diyor200/ecommerce/controllers"
	"github.com/diyor200/ecommerce/db"
	"github.com/diyor200/ecommerce/docs"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// @title Ecommerce api
// @description This API returns some data
// @version 1.0
// @host localhost:8000
// @BasePath /

func main() {
	const (
		port  = "8000"
		dburl = "postgres://postgres:2001@localhost:5432/ecommerce?sslmode=disable"
	)

	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	DB, err := sql.Open("postgres", dburl)
	if err != nil {
		panic(err)
	}

	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Muvaffaqiyatli ulandi!")

	err = db.CreateTables(DB)
	if err != nil {
		panic("Jadvallarni yaratishda xatolik yuzaga keldi!")
	}
	log.Println("Jadvalla yaratildi!")
	router := gin.Default()
	controllers.RegisterRoutes(router, DB)

	router.Run(":8000")

}
