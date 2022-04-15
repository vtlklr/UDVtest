package library

import (
	_ "github.com/jinzhu/gorm"
)

type Publishing struct {
	PublishingID uint   `gorm:"primary_key"`
	Name         string `json:"name"`
}
