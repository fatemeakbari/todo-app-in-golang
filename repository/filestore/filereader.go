package filestore

import (
	"os"
)

func readFileAsByte(filepath string) ([]byte, error) {

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var buffSize int64

	if fileStat, sErr := file.Stat(); sErr != nil {
		return nil, sErr
	} else {
		buffSize = fileStat.Size()
	}

	buff := make([]byte, buffSize)
	if _, rErr := file.Read(buff); rErr != nil {
		return nil, rErr
	}

	return buff, nil
}
