package filestore

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"todo/entity"
	"todo/logger"
	"todo/pkg"
	"todo/pkg/serializer"
	"todo/repository"
)

type userRepository struct {
	filepath string
	users    []entity.User

	hash       pkg.Hash
	serializer serializer.UserSerializer
}

func NewUserRepository(
	filepath string,
	hash pkg.Hash,
	userSerializer serializer.UserSerializer) (repository.UserRepository, error) {

	userRep := userRepository{
		filepath:   filepath,
		serializer: userSerializer,
		hash:       hash,
	}
	if err := userRep.load(); err != nil {
		return &userRep, err
	}

	return &userRep, nil
}

func (ur *userRepository) load() error {

	buff, err := readFileAsByte(ur.filepath)

	if err != nil {
		logger.LOGGER.Error(logger.RichError{
			MethodName: "load",
			Parent:     err,
			Message:    "problem in reading file " + ur.filepath},
		)

		return fmt.Errorf("problem in loading user storage")
	}
	rows := bytes.Split(buff, []byte("\n"))

	for _, row := range rows[:len(rows)-1] {
		var user entity.User
		//if string(row[0]) != "{" {
		//	continue
		//}
		if sErr := ur.serializer.Deserialize(row, &user); sErr != nil {
			logger.LOGGER.Error(logger.RichError{MethodName: "load", Parent: sErr})

			continue
		} else {

			ur.users = append(ur.users, user)
		}
	}
	logger.LOGGER.Info("user storage loaded successfully")
	return nil
}

func (ur *userRepository) Create(user entity.User) (entity.User, error) {

	file, err := os.OpenFile(ur.filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		logger.LOGGER.Error(logger.RichError{MethodName: "Create", Parent: err})

		return user, fmt.Errorf("problem in open storage file")
	}
	defer file.Close()

	user.ID = uint(len(ur.users)) + 1
	user.Password = ur.hash.Hash(user.Password)
	userByte, sErr := ur.serializer.Serialize(user)

	if sErr != nil {
		logger.LOGGER.Error(logger.RichError{MethodName: "Create", Parent: sErr})

		return user, fmt.Errorf("problem is save user")
	}

	_, wErr := file.Write(append(userByte, []byte("\n")...))

	if wErr != nil {
		logger.LOGGER.Error(logger.RichError{MethodName: "Create", Parent: wErr})

		return user, fmt.Errorf("problem is save user")
	}

	ur.users = append(ur.users, user)
	return user, nil
}

func (ur *userRepository) GetByEmailAndPassword(email, password string) (entity.User, error) {

	for _, user := range ur.users {
		if user.Email == email && user.Password == ur.hash.Hash(password) {
			return user, nil
		}
	}
	return entity.User{}, errors.New("user not found")
}
