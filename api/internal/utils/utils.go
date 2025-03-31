package utils

import (
	"reflect"
	"strings"
)

func If[T any](cond bool, a T, b T) T {
	if cond {
		return a
	} else {
		return b
	}
}

func NilSafeString(data *string) string {
	if data == nil {
		return ""
	}
	return *data
}

func MergeSlices[T any](slices ...[]T) []T {
	var result []T
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

func TrimStructStrings(v interface{}) {
	val := reflect.ValueOf(v).Elem() // Get the value of the struct pointer
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := val.Type().Field(i).Tag.Get("trim") // Get the trim tag

		// Check if the field is a string and has `trim:"true"`
		if field.Kind() == reflect.String && tag == "true" {
			field.SetString(strings.TrimSpace(field.String()))
		}
	}
}
