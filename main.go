package main

import (
	"crud_go/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.ConnectDB()

	if err != nil {
		log.Fatal("Can't connect to the database server")
	}

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/public", "./public")

	r.Run(":8080")
}
