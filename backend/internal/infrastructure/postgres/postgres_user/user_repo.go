package postgresuser

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repository"
	"fmt"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepo {
	return &GormUserRepository{db: db}
}

func (u *GormUserRepository) Create(user entities.User) error {
	newUser := UserToDB(user)

	if err := u.db.Create(newUser).Error; err != nil {
		return err
	}

	return nil
}

func (u *GormUserRepository) GetAll() ([]entities.User, error) {
	var userModels []UserModel

	if err := u.db.Find(&userModels).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}

	users := make([]entities.User, 0, len(userModels))
	for _, model := range userModels {

		user := entities.NewUser(
			model.Id,
			model.Username,
			model.Password,
			model.Role,
		)
		users = append(users, user)
	}

	return users, nil
}
