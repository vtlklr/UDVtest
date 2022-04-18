package repo

import (
	_ "github.com/jinzhu/gorm"
)

type User struct {
	UserID uint   `gorm:"primary_key"`
	Name   string `json:"name"`
}
