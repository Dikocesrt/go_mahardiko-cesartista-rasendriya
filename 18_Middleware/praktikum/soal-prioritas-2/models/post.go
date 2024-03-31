package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model `json:"-"`
	Id     int    `json:"id" form:"id" gorm:"primaryKey"`
	Judul  string `json:"title" form:"title"`
	Konten string `json:"content" form:"content"`
}