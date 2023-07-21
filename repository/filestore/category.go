package filestore

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"todo/entity"
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

	file, err := os.OpenFile(cr.filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("can not open file %s with error %w", cr.filepath, err)
	}

	var buffSize int64

	if fileStat, sErr := file.Stat(); sErr != nil {
		return fmt.Errorf("can not get stat of file %s with error %w", cr.filepath, sErr)
	} else {
		buffSize = fileStat.Size()
	}

	buff := make([]byte, buffSize)
	if _, rErr := file.Read(buff); rErr != nil {

		return fmt.Errorf("can not read file %s with error %w", cr.filepath, rErr)
	}

	rows := bytes.Split(buff, []byte("\n"))

	for _, row := range rows[:len(rows)-1] {
		var category entity.Category
		if string(row[0]) != "{" {
			continue
		}
		if sErr := cr.serializer.Deserialize(row, &category); sErr != nil {
			fmt.Println(sErr)
			continue
		} else {

			cr.categories = append(cr.categories, category)
		}
	}
	return nil
}
func (cr *categoryRepository) Create(category entity.Category) (entity.Category, error) {
	file, err := os.OpenFile(cr.filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		return category, fmt.Errorf("can not open file %s with error %w", cr.filepath, err)
	}
	defer file.Close()

	category.Id = uint(len(cr.categories)) + 1
	catByte, sErr := cr.serializer.Serialize(category)

	if sErr != nil {
		return category, fmt.Errorf("can not serialize category %w", sErr)
	}

	_, wErr := file.Write(append(catByte, []byte("\n")...))

	if wErr != nil {
		return category, fmt.Errorf("have a problem with writing  category byte, err %w", wErr)
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