package serializer

import (
	"todo/serializer/json"
	"todo/serializer/normal"
)

func GetUserSerializer(serializeMode string) UserSerializer {

	switch serializeMode {
	case "Normal":
		return normal.UserSerializer{}
	case "Json":
		return json.UserSerializer{}
	default:
		return json.UserSerializer{}
	}
}

func GetCategorySerializer(serializeMode string) CategorySerializer {

	switch serializeMode {
	case "Normal":
		return normal.CategorySerializer{}
	case "Json":
		return json.CategorySerializer{}
	default:
		return json.CategorySerializer{}
	}
}
