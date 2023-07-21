package normal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"todo/entity"
)

type CategorySerializer struct {
}

func (CategorySerializer) Serialize(c entity.Category) ([]byte, error) {
	return []byte(fmt.Sprintf("ID=%d, Title=%s, UserId=%d", c.ID, c.Title, c.UserId)), nil

}

func (CategorySerializer) Deserialize(cByte []byte, c *entity.Category) error {
	carr := strings.Split(strings.Trim(string(cByte), "\n"), ",")

	for _, item := range carr {
		field := strings.Split(item, "=")

		if len(field) == 2 {
			key, val := field[0], field[1]

			switch key {
			case "ID":
				id, _ := strconv.Atoi(val)
				c.ID = uint(id)
			case "Title":
				c.Title = val
			case "UserId":
				userId, _ := strconv.Atoi(val)
				c.UserId = uint(userId)

			}
		} else {

			return errors.New("your data is in wrong format")
		}
	}

	return nil
}
