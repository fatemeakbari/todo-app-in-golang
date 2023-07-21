package serializer

import (
	"todo/serializer/json"
	"todo/serializer/normal"
)

func GetUserSerializer(serializeMode string) UserSerializer {

	switch serializeMode {
	case "Normal":
		return normal.UserSerializer{}
	default:
		return json.UserSerializer{}
	}
}

func GetCategorySerializer(serializeMode string) CategorySerializer {

	switch serializeMode {
	case "Normal":
		return normal.CategorySerializer{}
	default:
		return json.CategorySerializer{}
	}
}

func GetTaskSerializer(serializeMode string) TaskSerializer {

	switch serializeMode {
	case "Normal":
		return normal.TaskSerializer{}
	default:
		return json.TaskSerializer{}
	}
}
