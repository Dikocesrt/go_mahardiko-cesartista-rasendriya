package user

import (
	"clean/constants"
	userEntitites "clean/entities/user"
	"clean/middleware"
)

type UserUseCase struct {
	repository userEntitites.RepositoryInterface
}

func NewUserUseCase(repository userEntitites.RepositoryInterface) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (userUseCase *UserUseCase) Register(user *userEntitites.User) (userEntitites.User, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return userEntitites.User{}, constants.ErrEmptyInput
	}
	
	newUser, err := userUseCase.repository.Register(user)
	if err != nil {
		return userEntitites.User{}, constants.ErrInsertDatabase
	}

	return newUser, nil
}

func (userUseCase *UserUseCase) Login(user *userEntitites.User) (userEntitites.User, error) {
	if user.Email == "" || user.Password == "" {
		return userEntitites.User{}, constants.ErrEmptyInput
	}

	userFromDb, err := userUseCase.repository.Login(user)
	if err != nil {
		return userEntitites.User{}, constants.ErrUserNotFound
	}

	token, _ := middleware.CreateToken(userFromDb.Id, userFromDb.Email)
	userFromDb.Token = token

	return userFromDb, nil
}