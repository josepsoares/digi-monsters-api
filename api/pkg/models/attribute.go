package models

import "gorm.io/gorm"

type Attribute struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}
