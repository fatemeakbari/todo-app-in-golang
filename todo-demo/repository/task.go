package repository

import (
	"time"
	"todo/entity"
)

type TaskRepository interface {
	Create(t *entity.Task) (entity.Task, error)
	GetAllUserTask(userId uint) ([]entity.Task, error)
	GetAllUserTaskByDueDate(dueDate time.Time, userId uint) ([]entity.Task, error)
	GetAllTodayUserTask(userId uint) ([]entity.Task, error)
	GetAllNonDoneUserTask(userId uint) ([]entity.Task, error)
}
