package main

import (
	"database/sql"
	"log"

	"github.com/diyor200/ecommerce/controllers"
	"github.com/diyor200/ecommerce/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	const (
		port  = "8000"
		dburl = "postgres://you_username:your_password@localhost:5432/dbname?sslmode=disable"
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
