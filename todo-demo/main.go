package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"todo/entity"
	"todo/pkg/sha256"
	"todo/repository"
	"todo/repository/filestore"
)

var (
	categoryStorage []entity.Category
	taskStorage     []entity.Task

	currentUser *entity.User

	reader = bufio.NewScanner(os.Stdin)

	userRepository repository.UserRepository
)

const (
	userStoragePath = "user.txt"
)

func main() {

	var err error
	if userRepository, err = filestore.NewUserRepository(
		userStoragePath,
		"Json",
		sha256.New()); err != nil {
		fmt.Println(err)
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

	fmt.Println("select a action")
	reader.Scan()
	command := reader.Text()

	switch command {

	case "task-list":
		_ = findAllUserTaskList()
	case "create-task":
		createTask()
	case "today-task-list":
	case "create-category":
		createCategory()
	case "category-list":
		_ = findAllUserCategoryList()
	case "exit":
	}
}

func createCategory() {
	fmt.Println("enter title")
	reader.Scan()
	title := reader.Text()

	fmt.Println("enter dueDate")
	reader.Scan()
	color := reader.Text()

	category := entity.Category{
		Id:     uint(len(categoryStorage)) + 1,
		Title:  title,
		Color:  color,
		UserId: currentUser.Id,
	}

	categoryStorage = append(categoryStorage, category)
}

func findAllUserCategoryList() []entity.Category {
	res := make([]entity.Category, 0)
	for _, cat := range categoryStorage {
		if cat.UserId == currentUser.Id {
			res = append(res, cat)
		}
	}

	fmt.Println(res)
	return res
}
func createTask() {

	fmt.Println("enter title")
	reader.Scan()
	title := reader.Text()

	fmt.Println("enter dueDate")
	reader.Scan()
	//dueDate := reader.Text()

	fmt.Println("enter categoryId")
	reader.Scan()
	categoryId, err := strconv.Atoi(reader.Text())

	if err != nil {
		return
	}

	if !isCategoryValid(categoryId) {
		fmt.Println("categoryId is wrong")
		return
	}

	task := entity.Task{
		Id:    uint(len(taskStorage)) + 1,
		Title: title,
		//DueDate: dueDate,
		IsDone: false,
		UserId: currentUser.Id}

	taskStorage = append(taskStorage, task)
	fmt.Println("task successfully created")
}

func isCategoryValid(categoryId int) bool {
	for _, cat := range categoryStorage {
		if cat.Id == uint(categoryId) && cat.UserId == currentUser.Id {
			return true
		}
	}
	return false
}
func findAllUserTaskList() []entity.Task {

	res := make([]entity.Task, 0)

	for _, task := range taskStorage {
		if task.UserId == currentUser.Id {
			res = append(res, task)
		}
	}

	fmt.Println(res)
	return res
}
