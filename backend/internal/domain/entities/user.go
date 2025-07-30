package entities

import "github.com/google/uuid"

type User struct {
	id       uuid.UUID
	username string
	password string
	role     string
}

func NewUser(id uuid.UUID, username, password, role string) User {
	return User{
		id:       id,
		username: username,
		password: password,
		role:     role,
	}
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Role() string {
	return u.role
}
