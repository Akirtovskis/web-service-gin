package models

import "gorm.io/gorm"

// album represents data about a record album.
type Album struct {
	gorm.Model

	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
