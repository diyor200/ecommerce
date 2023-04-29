package main

import (
	"database/sql"
	"log"

	"github.com/diyor200/ecommerce/controllers"
	"github.com/diyor200/ecommerce/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// @title           Ecommerce Web Api
// @version         1.0
// @description     Ecommerce web site api for exchanging information betwwn client and server

// @contact.name   Diyorbek Abdulaxatov
// @contact.url    https://t.me/Diyorbek01_31
// @contact.email  abdulaxatovdiyorbek40@gmail.com

// @host      localhost:8000
// @BasePath  /api

func main() {
	const (
		port  = "8000"
		dburl = "postgres://postgres:2001@localhost:5432/ecommerce?sslmode=disable"
	)

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
