package utils

import (
	"fmt"
	"golang.org/x/exp/slices"
	"reflect"
)

// Copy copies each value of source struct to dest struct by name,
// exclude those given in param exclude.
// source and dest must be struct pointer.
func Copy(sourcePtr any, destPtr any, exclude ...string) error {
	if reflect.TypeOf(sourcePtr).Kind() != reflect.Ptr ||
		reflect.TypeOf(destPtr).Kind() != reflect.Ptr {
		return fmt.Errorf("source and dest must be pointers")
	}

	sourceType := reflect.TypeOf(sourcePtr).Elem()
	destType := reflect.TypeOf(destPtr).Elem()
	if sourceType.Kind() != reflect.Struct ||
		destType.Kind() != reflect.Struct {
		return fmt.Errorf("source and dest must be structs")
	}

	source := reflect.ValueOf(sourcePtr).Elem()
	dest := reflect.ValueOf(destPtr).Elem()

	for i := 0; i < sourceType.NumField(); i++ {
		field := sourceType.Field(i)
		if slices.Contains(exclude, field.Name) {
			continue
		}
		sourceValue := source.FieldByName(field.Name)
		destValue := dest.FieldByName(field.Name)
		if destValue.IsValid() && destValue.CanSet() {
			destValue.Set(sourceValue)
		}
	}
	return nil
}
