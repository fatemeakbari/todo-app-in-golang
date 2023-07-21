package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"todo/cfg"
	"todo/entity"
	"todo/pkg/serializer"
	"todo/pkg/sha256"
	"todo/repository"
	"todo/repository/filestore"
)

var (
	currentUser *entity.User

	reader = bufio.NewScanner(os.Stdin)

	userRepository     repository.UserRepository
	categoryRepository repository.CategoryRepository
	taskRepository     repository.TaskRepository
)

func main() {

	var uErr error
	if userRepository, uErr = filestore.NewUserRepository(
		cfg.UserStoragePath,
		sha256.New(),
		serializer.GetUserSerializer(cfg.SerializerMode)); uErr != nil {
		fmt.Println(uErr)
		return
	}

	var cErr error
	if categoryRepository, cErr = filestore.NewCategoryRepository(
		cfg.CategoryStoragePath,
		serializer.GetCategorySerializer(cfg.SerializerMode)); cErr != nil {
		fmt.Println(cErr)
		return
	}

	var tErr error
	if taskRepository, tErr = filestore.NewTaskRepository(
		cfg.TaskStoragePath,
		serializer.GetTaskSerializer(cfg.SerializerMode)); tErr != nil {
		fmt.Println(tErr)
		return
	}

	for {
		if currentUser == nil {
			entranceProcess()
		}

		parseCommand()
	}
}

func entranceProcess() {

	for {
		fmt.Println("\nplease select \n1-login \n2-register \n3-exit\ncommand to continue")
		reader.Scan()
		command := reader.Text()

		if command == "register" {
			_ = register()
		} else if command == "login" {
			_ = login()
			if currentUser != nil {
				return
			}
		} else if command == "exit" {
			os.Exit(0)
		}

	}

}

func register() error {

	fmt.Println("enter name")
	reader.Scan()
	name := reader.Text()

	fmt.Println("enter email")
	reader.Scan()
	email := reader.Text()

	fmt.Println("enter password")
	reader.Scan()
	password := reader.Text()

	_, err := userRepository.Create(entity.User{
		Name:     name,
		Email:    email,
		Password: password,
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Welcome %s you successfully registered\n", name)
	}
	return err
}

func login() error {

	fmt.Println("enter email")
	reader.Scan()
	email := reader.Text()

	fmt.Println("enter password")
	reader.Scan()
	pass := reader.Text()

	user, err := userRepository.GetByEmailAndPassword(email, pass)

	if err != nil {
		fmt.Println("your email or password is wrong")
	} else {
		currentUser = &user
		fmt.Printf("Welcome %s you successfully logined\n", currentUser.Name)
	}

	return err
}

func parseCommand() {

	fmt.Println("select a action\n**you can enter action-list to see all actions")
	reader.Scan()
	command := reader.Text()

	switch command {

	case "action-list":
		actionList()
	case "task-list":
		fmt.Println(taskRepository.GetAllUserTask(currentUser.Id))
	case "create-task":
		createTask()
	case "today-task-list":
		fmt.Println(taskRepository.GetAllTodayDueDateUserTask(currentUser.Id))
	case "non-done-task-list":
		fmt.Println(taskRepository.GetAllNonDoneUserTask(currentUser.Id))
	case "create-category":
		createCategory()
	case "category-list":
		fmt.Println(categoryRepository.GetAllUserCategory(currentUser.Id))
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("your command is wrong please run action-list for more info")

	}
}

func actionList() {
	fmt.Println("1.create-category\n2.category-list\n3.create-task\n4.today-task-list\n5.non-don-task-list\n6.task-list\n7.exit")
}
func createCategory() {
	fmt.Println("enter title")
	reader.Scan()
	title := reader.Text()

	category := entity.Category{
		Title:  title,
		UserId: currentUser.Id,
	}

	cat, err := categoryRepository.Create(category)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("your category created successfully, category: ", cat)
	}
}

func createTask() {

	fmt.Println("enter title")
	reader.Scan()
	title := reader.Text()

	fmt.Println("enter dueTime\n**guid your time format must be same as\nYYYY-MM-DD HH:MM:SS for example 2006-01-02 15:04:05")

	reader.Scan()
	sDueDate := reader.Text()
	dueDate, err := time.Parse(cfg.TimestampFormat, sDueDate)
	if err != nil {
		fmt.Println("dueDate format is not correct")
		return
	}

	fmt.Println("enter categoryId")
	reader.Scan()
	categoryId, _ := strconv.Atoi(reader.Text())

	if !isCategoryValid(categoryId) {
		fmt.Println("categoryId is wrong")
		return
	}

	task := entity.Task{
		Title:      title,
		DueDate:    dueDate,
		IsDone:     false,
		CategoryId: uint(categoryId),
		UserId:     currentUser.Id}

	if _, err := taskRepository.Create(task); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("task successfully created")
	}
}

func isCategoryValid(categoryId int) bool {
	if _, err := categoryRepository.GetById(uint(categoryId)); err != nil {
		return false
	}
	return true
}
