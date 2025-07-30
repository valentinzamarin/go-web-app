package postgresuser

import "github.com/google/uuid"

type UserModel struct {
	Id       uuid.UUID `gorm:"primaryKey;type:uuid"`
	Username string    `gorm:"not null;unique"`
	Password string    `gorm:"not null"`
	Role     string    `gorm:"not null"`
}

func (UserModel) TableName() string {
	return "users"
}
