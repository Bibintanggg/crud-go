package main

import (
	"crud_go/database"
	"crud_go/internal/router"
	"log"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	r := router.SetupRoute(db)

	log.Println("🚀 Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server error:", err)
	}
}
