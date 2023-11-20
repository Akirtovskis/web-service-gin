package mocks

import (
	"errors"
	"strconv"
	"web-service-gin/models"
)

type MockDBService struct {
	SimulateError bool
}

// NewMockDBService creates a new instance of MockDBService
func NewMockDBService() *MockDBService {
	return &MockDBService{}
}

// GetAlbums returns a fixed set of albums for testing
func (m *MockDBService) GetAlbums() ([]models.Album, error) {

	// simulates error
	if m.SimulateError {
		return nil, errors.New("database error")
	}

	mockAlbums := []models.Album{
		{ID: 1, Title: "Mock Album 1", Artist: "Mock Artist", Price: 10.99},
		{ID: 2, Title: "Mock Album 2", Artist: "Mock Artist", Price: 12.99},
	}
	return mockAlbums, nil
}

// GetAlbumByID returns a mock album with the specified ID for testing
func (m *MockDBService) GetAlbumByID(id int) (models.Album, error) {
	// Create a mock album with the requested ID
	mockAlbum := models.Album{
		ID:     id,
		Title:  "Mock Album " + strconv.Itoa(id),
		Artist: "Mock Artist",
		Price:  10.99,
	}

	return mockAlbum, nil
}
