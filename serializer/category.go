package serializer

import "todo/entity"

type CategorySerializer interface {
	Serialize(c entity.Category) ([]byte, error)
	Deserialize(cByte []byte, u *entity.Category) error
}
