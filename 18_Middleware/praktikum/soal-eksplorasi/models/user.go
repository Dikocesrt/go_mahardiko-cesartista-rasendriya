package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Id       int    `json:"id" form:"id" gorm:"primaryKey"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	IsAdmin  int   `json:"isAdmin" form:"isAdmin"`
}

type UserResponse struct {
	gorm.Model `json:"-"`
	Id       int    `json:"id" form:"id" gorm:"primaryKey"`
	Email    string `json:"email" form:"email"`
	Token 	 string `json:"token" form:"token"`
}