package repo

import (
	_ "github.com/jinzhu/gorm"
)

type Library struct {
	LibBookID uint   `gorm:"primary_key"`
	BookID    uint   `json:"book_id"`
	Placement string `json:"placement"`
}
