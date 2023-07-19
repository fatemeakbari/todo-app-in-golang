package filestore

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"todo/entity"
	"todo/pkg"
	"todo/repository"
	"todo/serializer"
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

	file, err := os.OpenFile(ur.filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("can not open file %s with error %w", ur.filepath, err)
	}

	var buffSize int64

	if fileStat, sErr := file.Stat(); sErr != nil {
		return fmt.Errorf("can not get stat of file %s with error %w", ur.filepath, sErr)
	} else {
		buffSize = fileStat.Size()
	}

	buff := make([]byte, buffSize)
	if _, rErr := file.Read(buff); rErr != nil {

		return fmt.Errorf("can not read file %s with error %w", ur.filepath, rErr)
	}

	rows := bytes.Split(buff, []byte("\n"))

	for _, row := range rows[:len(rows)-1] {
		var user entity.User
		if string(row[0]) != "{" {
			continue
		}
		if sErr := ur.serializer.Deserialize(row, &user); sErr != nil {
			fmt.Println(sErr)
			continue
		} else {

			ur.users = append(ur.users, user)
		}
	}
	return nil
}

func (ur *userRepository) Create(user entity.User) (entity.User, error) {

	file, err := os.OpenFile(ur.filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		return user, fmt.Errorf("can not open file %s with error %w", ur.filepath, err)
	}
	defer file.Close()

	user.Id = uint(len(ur.users)) + 1
	user.Password = ur.hash.Hash(user.Password)
	userByte, sErr := ur.serializer.Serialize(user)

	if sErr != nil {
		return user, fmt.Errorf("can not serialize user %w", sErr)
	}

	_, wErr := file.Write(append(userByte, []byte("\n")...))

	if wErr != nil {
		return user, fmt.Errorf("have a problem with writing  user byte, err %w", wErr)
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
