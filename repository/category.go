package repository

import "todo/entity"

type CategoryRepository interface {
	Create(c entity.Category) (entity.Category, error)
	GetUserCategoryById(categoryId, userId uint) (entity.Category, error)
	GetAllUserCategory(userId uint) []entity.Category
}
