package main

import (
	"web-service-gin/config"
	"web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	router := gin.Default()
	routes.UseRoute(router)
	router.Run("localhost:8080")
}
