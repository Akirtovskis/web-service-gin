package album

import (
	"database/sql"
	"log"
	"net/http"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func (r *SqlAlbumRepository) GetAlbums(c *gin.Context) {
	var albums []models.Album

	rows, err := r.DB.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		log.Fatal(err)
	}

	// need to call defer rows.Close on query, but not on query row
	defer rows.Close()

	for rows.Next() {
		var alb models.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			log.Fatal(err)
		}
		albums = append(albums, alb)
	}

	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// GetAlbumById fetches an album by its ID
func (r *SqlAlbumRepository) GetAlbumById(c *gin.Context) {
	var album models.Album
	id := c.Param("id")

	// the $1 means to interpolate the first value after the query string
	row := r.DB.QueryRow("SELECT id, title, artist, price FROM albums WHERE id = $1", id)

	// Scan the result into the Album struct
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}
