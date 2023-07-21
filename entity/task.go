package entity

import (
	"fmt"
	"time"
	"todo/cfg"
)

type Task struct {
	ID      uint
	Title   string
	IsDone  bool
	DueDate time.Time

	UserId     uint
	CategoryId uint
}

func (t Task) String() string {
	return fmt.Sprintf("{ID: %d, Title: %s, IsDone: %t, DueDate: %s, CategoryId: %d", t.ID, t.Title, t.IsDone, t.DueDate.Format(cfg.TimestampFormat), t.CategoryId)
}
