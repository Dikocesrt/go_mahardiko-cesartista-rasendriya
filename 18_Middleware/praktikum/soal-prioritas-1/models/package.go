package models

import "gorm.io/gorm"

type Package struct {
	gorm.Model `json:"-"`
	Id             int     `json:"id" form:"id" gorm:"primaryKey"`
	Nama           string  `json:"name" form:"name"`
	NamaPengirim   string  `json:"sender" form:"sender"`
	NamaPenerima   string  `json:"receiver" form:"receiver"`
	LokasiPengirim string  `json:"sender_location" form:"sender_location"`
	LokasiPenerima string  `json:"sender_receiver" form:"sender_receiver"`
	Biaya          int     `json:"fee" form:"fee"`
	Berat          float64 `json:"weight" form:"weight"`
}