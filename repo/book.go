package repo

import (
	_ "github.com/jinzhu/gorm"
)

type Book struct {
	BookID       uint   `gorm:"primary_key"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Authors      string `json:"authors"`
	Year         uint   `json:"year"`
	Edition      string `json:"edition"`
	PublishingID int    `json:"publishing_id"`
}
