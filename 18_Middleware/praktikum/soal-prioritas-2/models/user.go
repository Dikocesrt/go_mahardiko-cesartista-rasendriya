package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Id         int    `json:"id" form:"id" gorm:"primaryKey"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
}

type UserResponse struct {
	Id    int    `json:"id" form:"id"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}