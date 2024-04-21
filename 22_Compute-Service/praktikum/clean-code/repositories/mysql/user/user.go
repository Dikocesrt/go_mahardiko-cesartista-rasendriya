package user

import (
	userEntities "clean/entities/user"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (userRepo *UserRepo) Register(user *userEntities.User) (userEntities.User, error) {
	userDb := FromUserEntitiesToUserDb(user)

	err := userRepo.DB.Create(&userDb).Error

	if err != nil {
		return userEntities.User{}, err
	}

	newUser := userDb.FromUserDbToUserEntities()

	return *newUser, nil
}

func (userRepo *UserRepo) Login(user *userEntities.User) (userEntities.User, error) {
	userDb := FromUserEntitiesToUserDb(user)

	err := userRepo.DB.Where("email = ? AND password = ?", userDb.Email, userDb.Password).First(&userDb).Error
	if err != nil {
		return userEntities.User{}, err
	}

	userFromDb := userDb.FromUserDbToUserEntities()

	return *userFromDb, nil
}