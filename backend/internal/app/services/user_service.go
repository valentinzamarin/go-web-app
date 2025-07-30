package services

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repository"
)

type UserService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) CreateNewUser(user entities.User) error {
	u.userRepo.Create(user)

	return nil
}

func (u *UserService) GetAllUsers() ([]entities.User, error) {

	users, err := u.userRepo.GetAll()
	if err != nil {
		return nil, nil
	}

	return users, nil
}
