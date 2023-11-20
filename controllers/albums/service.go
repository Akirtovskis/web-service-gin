package album

import (
	"database/sql"
	"web-service-gin/mocks"
	"web-service-gin/models"
)

// DEFINING INTERFACE FOR ALBUM REPOSITORY TO LATER REUSE IN DIFFERENT STRUCTS

// SQL REPOSITORY FOR PROD AND MOCK REPOSITORY FOR TESTS
type AlbumRepository interface {
	GetAlbums() ([]models.Album, error)
	GetAlbumByID(id int) (models.Album, error)
	CreateAlbum(album models.Album) error
	DeleteAlbum(id int) error
}

type SqlAlbumRepository struct {
	DB *sql.DB
}

// NewSqlAlbumRepository creates a new instance of SqlAlbumRepository
// kind of constructor
func NewSqlAlbumRepository(db *sql.DB) *SqlAlbumRepository {
	return &SqlAlbumRepository{DB: db}
}

type MockAlbumRepository struct {
	DBService *mocks.MockDBService
}

func NewMockAlbumRepository(dbService *mocks.MockDBService) *MockAlbumRepository {
	return &MockAlbumRepository{
		DBService: dbService,
	}
}

// Implementing the GetAlbums method to satisfy the AlbumRepository interface
func (repo *MockAlbumRepository) GetAlbums() ([]models.Album, error) {
	return repo.DBService.GetAlbums()
}

// Implementing the GetAlbumByID method to satisfy the AlbumRepository interface
func (repo *MockAlbumRepository) GetAlbumByID(id int) (models.Album, error) {
	return repo.DBService.GetAlbumByID(id)
}
