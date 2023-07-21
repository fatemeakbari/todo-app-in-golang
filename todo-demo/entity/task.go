package entity

import (
	"fmt"
	"time"
	"todo/cfg"
)

type Task struct {
	Id      uint
	Title   string
	IsDone  bool
	DueDate time.Time

	UserId     uint
	CategoryId uint
}

func (t Task) String() string {
	return fmt.Sprintf("{Id: %d, Title: %s, IsDone: %t, DueDate: %s, CategoryId: %d", t.Id, t.Title, t.IsDone, t.DueDate.Format(cfg.TIMESTAMP_FORMAT), t.CategoryId)
}
