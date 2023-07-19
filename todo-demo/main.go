package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"todo/entity"
)

var (
	userStorage     []entity.User
	categoryStorage []entity.Category
	taskStorage     []entity.Task

	currentUser *entity.User
	hash        = sha256.New()

	reader *bufio.Scanner = bufio.NewScanner(os.Stdin)
)

const (
	userStoragePath = "user.txt"
)

func main() {

	loadStorage()

	for {
		if currentUser == nil {
			entranceProcess()
		}

		parseCommand()
	}
}

func loadStorage() {

	loadUserStorage()

}

func addToUserStorage(user entity.User) error {

	file, err := os.OpenFile(userStoragePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	defer file.Close()

	if err != nil {
		fmt.Printf("have a problem with open the file %s, err= %v", userStoragePath, err)
		return err
	}

	userByte, err := json.Marshal(user)

	if err != nil {
		fmt.Printf("have a problem with marshaling user %s, err %v", user, err)
		return err
	}

	_, err = file.Write(append(userByte, []byte("\n")...))

	if err != nil {
		fmt.Printf("have a problem with writing  user byte, err %v", err)
		return err
	}

	return nil
}

func loadUserStorage() {
	file, err := os.OpenFile(userStoragePath, os.O_RDWR|os.O_APPEND, 0660)

	if err != nil {
		fmt.Printf("have a problem with open the file %s", userStoragePath)
	}

	buff := make([]byte, 1024)
	_, err = file.Read(buff)

	if err != nil {
		fmt.Printf("have a problem in read data from file %s err = %v", userStoragePath, err)
	}
	rows := bytes.Split(buff, []byte("\n"))

	for _, row := range rows {
		var user *entity.User
		if string(row[0]) != "{" {
			continue
		}
		err = json.Unmarshal(row, &user)

		if err != nil {
			fmt.Printf("have a problem in unmarshaling user data = %s, err = %v", row, err)
			continue
		}
		userStorage = append(userStorage, *user)
	}
}
func entranceProcess() {

	for {
		fmt.Println("\nplease select \n1-login \n2-register \n3-exit\ncommand to continue")
		reader.Scan()
		command := reader.Text()

		if command == "register" {
			register()
		} else if command == "login" {
			login()
			if currentUser != nil {
				return
			}
		} else if command == "exit" {
			os.Exit(0)
		}

	}

}

func register() {

	user := entity.User{Id: uint(len(userStorage)) + 1}

	fmt.Println("enter email")
	reader.Scan()
	user.Email = reader.Text()

	fmt.Println("enter password")
	reader.Scan()
	pass := reader.Text()
	user.Password = hex.EncodeToString(hash.Sum([]byte(pass)))

	err := addToUserStorage(user)
	if err != nil {
		return
	}

	userStorage = append(userStorage, user)

	fmt.Println("you successfully registered, please login now")
}
func login() {

	fmt.Println("enter email")
	reader.Scan()
	email := reader.Text()

	fmt.Println("enter password")
	reader.Scan()
	pass := reader.Text()

	currentUser = findUser(email, pass)

	if currentUser == nil {
		fmt.Println("your email or password is wrong please try again")
	} else {
		fmt.Println("you are login successfully")
	}
}
func findUser(email, password string) *entity.User {
	for _, user := range userStorage {

		if user.Email == email && user.Password == hex.EncodeToString(hash.Sum([]byte(password))) {
			return &user
		}
	}
	return nil
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
	dueDate := reader.Text()

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
		Id:      uint(len(taskStorage)) + 1,
		Title:   title,
		DueDate: dueDate,
		IsDone:  false,
		UserId:  currentUser.Id}

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
