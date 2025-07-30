package interfaces

import "backend/internal/domain/entities"

type UserService interface {
	CreateNewUser(user entities.User) error
	GetAllUsers() ([]entities.User, error)
}
