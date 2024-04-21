package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id int
	Name string
	Email string
	Password string
	Token string
}

type RepositoryInterface interface {
	Register(user *User) (User, error)
	Login(user *User) (User, error)
}

type UseCaseInterface interface {
	Register(user *User) (User, error)
	Login(user *User) (User, error)
}