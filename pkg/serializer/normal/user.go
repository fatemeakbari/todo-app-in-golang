package normal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"todo/entity"
)

type UserSerializer struct {
}

func (s UserSerializer) Serialize(u entity.User) ([]byte, error) {
	return []byte(fmt.Sprintf("ID=%d, Name=%s, Email=%s, Password=%s", u.ID, u.Name, u.Email, u.Password)), nil
}

func (s UserSerializer) Deserialize(userByte []byte, u *entity.User) error {

	uarr := strings.Split(strings.Trim(string(userByte), "\n"), ",")

	for _, item := range uarr {
		field := strings.Split(item, "=")

		if len(field) == 2 {
			key, val := field[0], field[1]

			switch key {
			case "ID":
				if id, err := strconv.Atoi(val); err != nil {

					return fmt.Errorf(`user ID %v is in wrong format, err: %w`, val, err)
				} else {
					u.ID = uint(id)
				}
			case "Name":
				u.Name = val
			case "Email":
				u.Email = val
			case "Password":
				u.Password = val

			}
		} else {

			return errors.New("your data is in wrong format")
		}
	}

	return nil
}
