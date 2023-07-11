package utils

import "reflect"

func IsStructEmpty(data interface{}) bool {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Interface() != reflect.Zero(field.Type()).Interface() {
			return false
		}
	}

	return true
}
