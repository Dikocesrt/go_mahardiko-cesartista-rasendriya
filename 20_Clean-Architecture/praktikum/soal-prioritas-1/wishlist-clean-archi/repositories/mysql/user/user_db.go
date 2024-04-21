package user

import (
	userEntities "clean/entities/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}

func FromUserEntitiesToUserDb(user *userEntities.User) *User {
	return &User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func (user *User) FromUserDbToUserEntities() *userEntities.User {
	return &userEntities.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
	}
}