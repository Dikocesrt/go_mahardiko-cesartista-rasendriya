package model

type Comment struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	IsiKomentar string `json:"content"`
	PostinganId int    `gorm:"index"`
}