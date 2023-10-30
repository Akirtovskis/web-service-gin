package routes

import (
	"net/http"
	"time"
	"web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
// var albums = []common.Album{
// 	{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok", "timestap": time.Now().Format(time.RFC3339)})
}

func UseRoute(router *gin.Engine) {
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/albums", controllers.CreateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)
	router.GET("/albums/:id", controllers.GetAlbumById)
	router.GET("/health", getHealth)
	router.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid path"})
	})
	router.Run("localhost:8080")
}
