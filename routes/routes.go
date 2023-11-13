package routes

import (
	"database/sql"
	"net/http"
	"time"

	album "web-service-gin/controllers/albums"

	"github.com/gin-gonic/gin"
)

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok", "timestap": time.Now().Format(time.RFC3339)})
}

func UseRoute(router *gin.Engine, db *sql.DB) {

	// Set up the album service.
	albumService := album.NewAlbumService(db)

	// v important to get more complex routes in the beginning so they are not overwritten by simpler routes
	router.GET("/albums/:id", func(c *gin.Context) { albumService.GetAlbumById(c) })
	router.GET("/albums", func(c *gin.Context) { albumService.GetAlbums(c) })
	router.POST("/albums", func(c *gin.Context) { albumService.CreateAlbum(c) })
	router.DELETE("/albums/:id", func(c *gin.Context) { albumService.DeleteAlbum(c) })

	router.GET("/health", getHealth)
	router.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid path"})
	})

}
