package entity

import "fmt"

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

func (u User) String() string {
	return fmt.Sprintf("{ID: %d, Name: %s, Email: %s}", u.ID, u.Name, u.Email)
}
