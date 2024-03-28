package model

type Post struct {
	Id       int       `json:"id" gorm:"primaryKey"`
	Judul    string    `json:"title"`
	Konten   string    `json:"content"`
	Comments []Comment `json:"comments" gorm:"-"`
}