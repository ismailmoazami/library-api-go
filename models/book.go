package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
