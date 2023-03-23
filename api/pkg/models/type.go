package models

import "gorm.io/gorm"

// Book model
type Type struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}
