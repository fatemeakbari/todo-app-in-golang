package serializer

import (
	json2 "todo/pkg/serializer/json"
	normal2 "todo/pkg/serializer/normal"
)

func GetUserSerializer(serializeMode string) UserSerializer {

	switch serializeMode {
	case "Normal":
		return normal2.UserSerializer{}
	default:
		return json2.UserSerializer{}
	}
}

func GetCategorySerializer(serializeMode string) CategorySerializer {

	switch serializeMode {
	case "Normal":
		return normal2.CategorySerializer{}
	default:
		return json2.CategorySerializer{}
	}
}

func GetTaskSerializer(serializeMode string) TaskSerializer {

	switch serializeMode {
	case "Normal":
		return normal2.TaskSerializer{}
	default:
		return json2.TaskSerializer{}
	}
}
