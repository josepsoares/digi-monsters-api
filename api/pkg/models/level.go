package models

import "gorm.io/gorm"

type Level struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}
