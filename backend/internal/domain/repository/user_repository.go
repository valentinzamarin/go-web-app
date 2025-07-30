package repository

import "backend/internal/domain/entities"

type UserRepo interface {
	Create(user entities.User) error
	GetAll() ([]entities.User, error)
}
