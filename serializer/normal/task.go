package normal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"todo/cfg"
	"todo/entity"
)

type TaskSerializer struct {
}

func (s TaskSerializer) Serialize(t entity.Task) ([]byte, error) {
	return []byte(fmt.Sprintf("Id=%d, Title=%s, IsDone=%t, DueDate=%s, UserId=%d, CategoryId=%d", t.Id, t.Title, t.IsDone, t.DueDate.Format(cfg.TIMESTAMP_FORMAT), t.UserId, t.CategoryId)), nil
}

func (s TaskSerializer) Deserialize(tByte []byte, t *entity.Task) error {

	tarr := strings.Split(strings.Trim(string(tByte), "\n"), ",")

	for _, item := range tarr {
		field := strings.Split(item, "=")

		if len(field) == 2 {
			key, val := field[0], field[1]

			switch key {
			case "Id":
				id, _ := strconv.Atoi(val)
				t.Id = uint(id)
			case "Title":
				t.Title = val
			case "IsDone":
				t.IsDone, _ = strconv.ParseBool(val)
			case "DueDate":
				t.DueDate, _ = time.Parse(cfg.TIMESTAMP_FORMAT, val)
			case "UserId":
				userId, _ := strconv.Atoi(val)
				t.UserId = uint(userId)
			case "CategoryId":
				cId, _ := strconv.Atoi(val)
				t.CategoryId = uint(cId)

			}
		} else {

			return errors.New("your data is in wrong format")
		}
	}

	return nil
}
