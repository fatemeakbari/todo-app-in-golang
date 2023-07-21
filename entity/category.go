package entity

import "fmt"

type Category struct {
	ID    uint
	Title string

	UserId uint
}

func (c Category) String() string {
	return fmt.Sprintf("{ID: %d, Title: %s}", c.ID, c.Title)
}
