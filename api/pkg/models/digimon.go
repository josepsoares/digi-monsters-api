package models

import "gorm.io/gorm"

type Digimon struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}
