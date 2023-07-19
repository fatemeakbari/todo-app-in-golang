package json

import (
	"encoding/json"
	"todo/entity"
)

type UserSerializer struct {
}

func (s UserSerializer) Serialize(u entity.User) ([]byte, error) {
	return json.Marshal(u)
}

func (s UserSerializer) Deserialize(userByte []byte, u *entity.User) error {
	return json.Unmarshal(userByte, u)
}
