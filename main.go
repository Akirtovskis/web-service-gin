package main

import (
	"log"
	"web-service-gin/config"
	"web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	router := gin.Default()
	routes.UseRoute(router, db)
	router.Run(":8080")
}
