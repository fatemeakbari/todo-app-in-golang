package repository

import "todo/entity"

type CategoryRepository interface {
	Create(c *entity.Category) (entity.Category, error)
	GetById(categoryId uint) (entity.Category, error)
	GetAllUserCategory(userId uint) ([]entity.Category, error)
}
