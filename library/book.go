package library

import (
	_ "github.com/jinzhu/gorm"
)

type Book struct {
	BookID       uint   `gorm:"primary_key"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Authors      string `json:"authors"`
	Year         int    `json:"year"`
	Edition      string `json:"edition"`
	PublishingID int    `json:"publishing_id"`
}
