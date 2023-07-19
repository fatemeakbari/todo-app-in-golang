package repository

import (
	"todo/entity"
)

type UserRepository interface {
	Create(u entity.User) (entity.User, error)
	GetByEmailAndPassword(email, password string) (entity.User, error)
}
