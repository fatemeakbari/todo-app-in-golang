package serializer

import "todo/entity"

type TaskSerializer interface {
	Serialize(t entity.Task) ([]byte, error)
	Deserialize(tByte []byte, u *entity.Task) error
}
