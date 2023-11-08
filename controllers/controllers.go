package controllers

import (
	"net/http"
	"web-service-gin/config"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	albumsFromDb := []models.Album{}
	config.DB.Find(&albumsFromDb)
	c.IndentedJSON(http.StatusOK, albumsFromDb)
}

// postAlbums adds an album from JSON received in the request body.
func CreateAlbum(c *gin.Context) {
	var newAlbum models.Album
	c.BindJSON(&newAlbum)
	config.DB.Create(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// delete album from db by id
func DeleteAlbum(c *gin.Context) {
	var albumToDelete models.Album
	config.DB.Where("id = ?", c.Param("id")).Delete(&albumToDelete)

	response := gin.H{
		"message": "Album deleted",
		"albumID": c.Param("id"),
	}

	c.IndentedJSON(http.StatusOK, response)
}

// get album by id
func GetAlbumById(c *gin.Context) {
	var album models.Album
	config.DB.Where("id = ?", c.Param("id")).Find(&album)
	c.IndentedJSON(http.StatusOK, album)
}
