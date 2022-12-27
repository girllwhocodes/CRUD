package models

import "gorm.io/gorm"

type Education struct {
	gorm.Model
	Title   string `json:"title"`
	Country string `json:"country"`
	Date    string `json:"date"`
}
