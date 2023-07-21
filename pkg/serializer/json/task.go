package json

import (
	"encoding/json"
	"todo/entity"
)

type TaskSerializer struct {
}

func (s TaskSerializer) Serialize(t entity.Task) ([]byte, error) {
	return json.Marshal(t)
}

func (s TaskSerializer) Deserialize(tByte []byte, t *entity.Task) error {
	return json.Unmarshal(tByte, t)
}
