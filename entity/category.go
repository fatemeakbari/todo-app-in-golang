package entity

import "fmt"

type Category struct {
	Id    uint
	Title string

	UserId uint
}

func (c Category) String() string {
	return fmt.Sprintf("{Id: %d, Title: %s}", c.Id, c.Title)
}
