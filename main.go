package main

import (
	"database/sql"
	"fmt"

	"github.com/diyor200/ecommerce/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	const (
		port  = "8000"
		dburl = "postgres://postgres:2001@localhost:5432/testgo?sslmode=disable"
	)

	db, err := sql.Open("postgres", dburl)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	router := gin.Default()
	controllers.RegisterRoutes(router, db)

	router.Run(":8000")

}
