package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Title       string
	Description string
	Price       float64
	UserID      uint
}
