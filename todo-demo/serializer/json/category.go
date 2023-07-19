package json

import (
	"encoding/json"
	"todo/entity"
)

type CategorySerializer struct {
}

func (CategorySerializer) Serialize(c entity.Category) ([]byte, error) {
	return json.Marshal(c)
}

func (CategorySerializer) Deserialize(cByte []byte, c *entity.Category) error {
	return json.Unmarshal(cByte, c)
}
