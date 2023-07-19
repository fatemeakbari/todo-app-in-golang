package entity

import "time"

type Task struct {
	Id      uint
	Title   string
	IsDone  bool
	DueDate time.Time

	UserId     uint
	CategoryId uint
}
