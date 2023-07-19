package serializer

import (
	"todo/entity"
)

type UserSerializer interface {
	Serialize(u entity.User) ([]byte, error)
	Deserialize(userByte []byte, u *entity.User) error
}
