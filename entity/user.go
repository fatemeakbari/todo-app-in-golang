package entity

import "fmt"

type User struct {
	Id       uint
	Name     string
	Email    string
	Password string
}

func (u User) String() string {
	return fmt.Sprintf("{Id: %d, Name: %s, Email: %s}", u.Id, u.Name, u.Email)
}
