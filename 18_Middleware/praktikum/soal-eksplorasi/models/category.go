package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model `json:"-"`
	Id   int    `json:"id" form:"id" gorm:"primaryKey"`
	Nama string `json:"name" form:"name"`
}