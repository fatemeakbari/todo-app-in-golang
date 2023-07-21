package repository

import (
	"time"
	"todo/entity"
)

type TaskRepository interface {
	Create(t entity.Task) (entity.Task, error)
	GetAllUserTask(userId uint) []entity.Task
	GetAllUserTaskByDueDate(dueDate time.Time, userId uint) []entity.Task
	GetAllTodayDueDateUserTask(userId uint) []entity.Task
	GetAllNonDoneUserTask(userId uint) []entity.Task
}
