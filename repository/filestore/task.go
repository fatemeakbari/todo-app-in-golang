package filestore

import (
	"bytes"
	"fmt"
	"os"
	"time"
	"todo/cfg"
	"todo/entity"
	"todo/repository"
	"todo/serializer"
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

	file, err := os.OpenFile(tr.filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("can not open file %s with error %w", tr.filepath, err)
	}

	var buffSize int64

	if fileStat, sErr := file.Stat(); sErr != nil {
		return fmt.Errorf("can not get stat of file %s with error %w", tr.filepath, sErr)
	} else {
		buffSize = fileStat.Size()
	}

	buff := make([]byte, buffSize)
	if _, rErr := file.Read(buff); rErr != nil {

		return fmt.Errorf("can not read file %s with error %w", tr.filepath, rErr)
	}

	rows := bytes.Split(buff, []byte("\n"))

	for _, row := range rows[:len(rows)-1] {
		var task entity.Task

		if sErr := tr.serializer.Deserialize(row, &task); sErr != nil {
			fmt.Println(sErr)
			continue
		} else {

			tr.tasks = append(tr.tasks, task)
		}
	}
	return nil
}

func (tr *taskRepository) Create(task entity.Task) (entity.Task, error) {
	file, err := os.OpenFile(tr.filepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return task, fmt.Errorf("can not open file %s, err: %w", tr.filepath, err)
	}
	defer file.Close()

	task.Id = uint(len(tr.tasks)) + 1
	tByte, sErr := tr.serializer.Serialize(task)
	if sErr != nil {
		return task, fmt.Errorf("can not serialize task, err: %w", sErr)
	}

	if _, wErr := file.Write(append(tByte, []byte("\n")...)); wErr != nil {
		return task, fmt.Errorf("can not write task, err: %w", wErr)
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
