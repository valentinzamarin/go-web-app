package postgresuser

import "backend/internal/domain/entities"

func UserToDB(user entities.User) *UserModel {
	return &UserModel{
		Id:       user.ID(),
		Username: user.Username(),
		Password: user.Password(),
		Role:     user.Role(),
	}
}

func DBToUser(model UserModel) entities.User {
	return entities.NewUser(
		model.Id,
		model.Username,
		model.Password,
		model.Role,
	)
}
