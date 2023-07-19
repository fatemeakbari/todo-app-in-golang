package repository

import (
	"os/user"
	"todo/entity"
)

type UserRepository interface {
	Create(u *user.User) (entity.User, error)
	GetByEmailAndPassword(email, password string) (entity.User, error)
}
