package routes

import (
	"database/sql"
	"net/http"
	"time"

	album "web-service-gin/controllers/albums"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
// var albums = []models.Album{
// 	{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok", "timestap": time.Now().Format(time.RFC3339)})
}

func UseRoute(router *gin.Engine, db *sql.DB) {
	// Set up the album service.
	albumService := album.NewAlbumService(db)

	router.GET("/albums", func(c *gin.Context) { albumService.GetAlbums(c, db) })
	// TODO THESE ROUTES ARE NOT IMPLEMENTED YET
	// ------------------------------------------------
	// router.POST("/albums", albumService.CreateAlbum)
	// router.DELETE("/albums/:id", albumService.DeleteAlbum)
	// router.GET("/albums/:id", albumService.GetAlbumById)
	router.GET("/health", getHealth)
	router.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid path"})
	})
	router.Run(":8080")
}
