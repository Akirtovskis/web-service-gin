package album

import (
	"database/sql"
	"log"
	"net/http"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func (s *AlbumService) GetAlbums(c *gin.Context, db *sql.DB) {
	var albums []models.Album

	rows, err := db.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		log.Fatal(err)
	}
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

// old ORM integration should be migrated to SQL
// ------------------------------------------------
// // get album by id
// func GetAlbumById(c *gin.Context) {
// 	var album models.Album
// 	config.DB.Where("id = ?", c.Param("id")).Find(&album)
// 	c.IndentedJSON(http.StatusOK, album)
// }
