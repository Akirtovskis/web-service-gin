package album

import (
	"database/sql"
)

// defining a struct which is essentialy like defining a class in JavaScript
type AlbumService struct {
	DB *sql.DB
}

// defining a method on the AlbumService struct is like defining a method on a class in JavaScript
func NewAlbumService(db *sql.DB) *AlbumService {
	return &AlbumService{DB: db}
}
