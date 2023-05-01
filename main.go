package main

import (
	"database/sql"
	"log"

	"github.com/diyor200/ecommerce/controllers"
	"github.com/diyor200/ecommerce/db"
	"github.com/diyor200/ecommerce/docs"
	_ "github.com/diyor200/ecommerce/docs"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

//	@title			Ecommerce Web api
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

func main() {

	const (
		dburl = "postgres://postgres:2001@localhost:5432/ecommerce?sslmode=disable"
	)

	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"

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
