package filestore

import (
	"bytes"
	"fmt"
	"os"
	"time"
	"todo/cfg"
	"todo/entity"
	"todo/logger"
	"todo/pkg/serializer"
	"todo/repository"
)

type taskRepository struct {
	filepath string
	tasks    []entity.Task

	serializer serializer.TaskSerializer
}

func NewTaskRepository(filepath string, taskSerializer serializer.TaskSerializer) (repository.TaskRepository, error) {
	taskRep := &taskRepository{
		filepath:   filepath,
		serializer: taskSerializer,
	}

	if err := taskRep.load(); err != nil {
		return &taskRepository{}, err
	}
	return taskRep, nil
}

func (tr *taskRepository) load() error {

	buff, err := readFileAsByte(tr.filepath)

	if err != nil {
		logger.LOGGER.Error(logger.RichError{
			MethodName: "load",
			Parent:     err,
			Message:    "problem in reading file " + tr.filepath},
		)

		return fmt.Errorf("problem in loading user storage")
	}

	rows := bytes.Split(buff, []byte("\n"))

	for _, row := range rows {
		var task entity.Task

		if sErr := tr.serializer.Deserialize(row, &task); sErr != nil {
			logger.LOGGER.Error(logger.RichError{MethodName: "load", Parent: sErr})

			continue
		} else {

			tr.tasks = append(tr.tasks, task)
		}
	}

	logger.LOGGER.Info("task storage loaded successfully")

	return nil
}

func (tr *taskRepository) Create(task entity.Task) (entity.Task, error) {
	file, err := os.OpenFile(tr.filepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		logger.LOGGER.Error(logger.RichError{MethodName: "Create", Parent: err})

		return task, fmt.Errorf("problem in open storage file")
	}
	defer file.Close()

	task.Id = uint(len(tr.tasks)) + 1
	tByte, sErr := tr.serializer.Serialize(task)
	if sErr != nil {
		logger.LOGGER.Error(logger.RichError{MethodName: "Create", Parent: sErr})

		return task, fmt.Errorf("problem in serialzie task")
	}

	if _, wErr := file.Write(append(tByte, []byte("\n")...)); wErr != nil {
		logger.LOGGER.Error(logger.RichError{MethodName: "Create", Parent: wErr})

		return task, fmt.Errorf("problem in write task")
	}

	tr.tasks = append(tr.tasks, task)

	return task, nil

}
func (tr *taskRepository) GetAllUserTask(userId uint) []entity.Task {

	var res []entity.Task

	for _, task := range tr.tasks {
		if task.UserId == userId {
			res = append(res, task)
		}
	}

	return res
}
func (tr *taskRepository) GetAllUserTaskByDueDate(dueDate time.Time, userId uint) []entity.Task {
	var res []entity.Task

	for _, task := range tr.tasks {

		if dueDate.Format(cfg.DateFormat) == task.DueDate.Format(cfg.DateFormat) &&
			task.UserId == userId {
			res = append(res, task)
		}
	}

	return res
}
func (tr *taskRepository) GetAllTodayDueDateUserTask(userId uint) []entity.Task {

	return tr.GetAllUserTaskByDueDate(time.Now(), userId)
}

func (tr *taskRepository) GetAllNonDoneUserTask(userId uint) []entity.Task {
	var res []entity.Task

	for _, task := range tr.tasks {
		if !task.IsDone && task.UserId == userId {
			res = append(res, task)
		}
	}

	return res
}
