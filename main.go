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

	// Log registered routes before starting the server
	// --- good for debugging when adding new routes
	// for _, routeInfo := range router.Routes() {
	// 	log.Printf("method: %s, path: %s\n", routeInfo.Method, routeInfo.Path)
	// }

	router.Run(":8080")
}
