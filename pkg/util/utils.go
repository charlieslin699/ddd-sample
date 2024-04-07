package util

import (
	"encoding/json"
	"reflect"
)

func GetStructName(s any) string {
	return reflect.TypeOf(s).Name()
}

func JSONMarshal[T any](data T) (jsonString string, err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func JSONUnmarshal[T any](jsonString string) (data T, err error) {
	err = json.Unmarshal([]byte(jsonString), &data)
	return
}

func IndexOf[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}

	return -1
}
