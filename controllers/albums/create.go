package album

import (
	"net/http"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func (r *SqlAlbumRepository) CreateAlbum(c *gin.Context) {
	var newAlbum models.Album

	// Bind the incoming JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Prepare and execute the SQL statement with RETURNING id
	query := "INSERT INTO albums (title, artist, price) VALUES ($1, $2, $3) RETURNING id"
	err := r.DB.QueryRow(query, newAlbum.Title, newAlbum.Artist, newAlbum.Price).Scan(&newAlbum.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "database error from preparing and executing"})
		return
	}

	// Respond with the created album and a success message
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "album has been successfully added",
		"album":   newAlbum,
	})
}
