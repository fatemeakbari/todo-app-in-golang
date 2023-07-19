package entity

type Task struct {
	Id      uint
	Title   string
	IsDone  bool
	DueDate string

	UserId     uint
	CategoryId uint
}
