package postgres

import (
	postgresuser "backend/internal/infrastructure/postgres/postgres_user"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&postgresuser.UserModel{},
	)
}
