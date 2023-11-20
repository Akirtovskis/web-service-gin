package album

import (
	"errors"
	"testing"
	"web-service-gin/mocks"
	"web-service-gin/models"

	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	// Create a mock instance of the DB service
	mockDBService := mocks.NewMockDBService()

	// Create an instance of MockAlbumRepository with the mocked service
	mockRepo := NewMockAlbumRepository(mockDBService)

	// Call the method we want to test
	albums, err := mockRepo.GetAlbums()

	// Define what we expect to receive
	expectedAlbums := []models.Album{
		{ID: 1, Title: "Mock Album 1", Artist: "Mock Artist", Price: 10.99},
		{ID: 2, Title: "Mock Album 2", Artist: "Mock Artist", Price: 12.99},
	}

	// Assertions to verify the behavior and the results
	assert.NoError(t, err)
	assert.NotNil(t, albums)
	assert.Equal(t, expectedAlbums, albums)
}

func TestGetAlbumsWithError(t *testing.T) {
	mockDB := mocks.NewMockDBService()
	mockDB.SimulateError = true

	// Call GetAlbums and expect an error
	_, err := mockDB.GetAlbums()

	if err == nil {
		t.Errorf("Expected an error but got none")
	}

	assert.Equal(t, err, errors.New("database error"))
}

func TestGetAlbumsById(t *testing.T) {
	// Create a mock instance of the DB service
	mockDBService := mocks.NewMockDBService()

	// Create an instance of MockAlbumRepository with the mocked service
	mockRepo := NewMockAlbumRepository(mockDBService)

	// Call the method we want to test
	album, err := mockRepo.GetAlbumByID(1)

	// Define what we expect to receive
	expectedAlbum := models.Album{ID: 1, Title: "Mock Album 1", Artist: "Mock Artist", Price: 10.99}

	// Assertions to verify the behavior and the results
	assert.NoError(t, err)
	assert.NotNil(t, album)
	assert.Equal(t, expectedAlbum, album)
}
