package filestore

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"todo/entity"
	"todo/logger"
	"todo/repository"
	"todo/serializer"
)

type categoryRepository struct {
	filepath   string
	categories []entity.Category

	serializer serializer.CategorySerializer
}

func NewCategoryRepository(filepath string, serializer serializer.CategorySerializer) (repository.CategoryRepository, error) {
	categoryRep := &categoryRepository{
		filepath:   filepath,
		serializer: serializer,
	}
	if err := categoryRep.load(); err != nil {
		return categoryRep, err
	}
	return categoryRep, nil
}

func (cr *categoryRepository) load() error {

	buff, err := readFileAsByte(cr.filepath)

	if err != nil {
		logger.LOGGER.Error(logger.RichError{
			MethodName: "load",
			Parent:     err,
			Message:    "problem in reading file " + cr.filepath},
		)

		return fmt.Errorf("problem in loading user storage")
	}

	rows := bytes.Split(buff, []byte("\n"))

	for _, row := range rows[:len(rows)-1] {
		var category entity.Category
		//if string(row[0]) != "{" {
		//	continue
		//}
		if sErr := cr.serializer.Deserialize(row, &category); sErr != nil {
			logger.LOGGER.Error(logger.RichError{MethodName: "load", Parent: sErr})

			continue
		} else {

			cr.categories = append(cr.categories, category)
		}
	}

	logger.LOGGER.Info("category storage loaded successfully")

	return nil
}
func (cr *categoryRepository) Create(category entity.Category) (entity.Category, error) {
	file, err := os.OpenFile(cr.filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		logger.LOGGER.Error(logger.RichError{MethodName: "Create", Parent: err})

		return category, fmt.Errorf("problem in open storage file")
	}
	defer file.Close()

	category.Id = uint(len(cr.categories)) + 1
	catByte, sErr := cr.serializer.Serialize(category)

	if sErr != nil {
		logger.LOGGER.Error(logger.RichError{MethodName: "Create", Parent: sErr})

		return category, fmt.Errorf("problem in serialize category")
	}

	_, wErr := file.Write(append(catByte, []byte("\n")...))

	if wErr != nil {
		logger.LOGGER.Error(logger.RichError{MethodName: "Create", Parent: wErr})

		return category, fmt.Errorf("problem in write category")
	}

	cr.categories = append(cr.categories, category)
	return category, nil
}

func (cr *categoryRepository) GetById(categoryId uint) (entity.Category, error) {

	for _, cat := range cr.categories {
		if cat.Id == categoryId {
			return cat, nil
		}
	}

	return entity.Category{}, errors.New("category not found")
}

func (cr *categoryRepository) GetAllUserCategory(userId uint) []entity.Category {
	var res []entity.Category

	for _, cat := range cr.categories {
		if cat.UserId == userId {
			res = append(res, cat)
		}
	}
	return res
}
