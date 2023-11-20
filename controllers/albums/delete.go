package album

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *SqlAlbumRepository) DeleteAlbum(c *gin.Context) {
	// Convert the 'id' URL parameter to an integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid album ID"})
		return
	}

	// Prepare the SQL statement
	stmt, err := r.DB.Prepare("DELETE FROM albums WHERE id = $1")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// need to close the prepare statemnt after execution
	defer stmt.Close()

	// Execute the statement with the album ID
	_, err = stmt.Exec(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// Create a response message
	response := gin.H{
		"message": "Album deleted",
		"albumID": id,
	}

	// Respond with the message
	c.IndentedJSON(http.StatusOK, response)
}
