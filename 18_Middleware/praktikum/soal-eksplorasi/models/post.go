package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model `json:"-"`
	Id         int      `json:"id" form:"id" gorm:"primaryKey"`
	Judul      string   `json:"title" form:"title"`
	Konten     string   `json:"content" form:"content"`
	CategoryId int      `json:"category_id" form:"category_id"`
	Category   Category `gorm:"primaryKey:categoryId"`
}

type PostResponse struct {
	Id       int      `json:"id"`
	Judul    string   `json:"title"`
	Konten   string   `json:"content"`
	Category Category `json:"category"`
}