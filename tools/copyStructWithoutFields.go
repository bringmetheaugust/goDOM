package tools

import (
	"fmt"
	"reflect"
)

// Copy struct without passed fields.
func СopyStructWithoutFields[T any](input T, fieldsToIgnore []string) (T, error) {
	v := reflect.ValueOf(input)

	if v.Kind() != reflect.Struct {
		return input, fmt.Errorf("input is not a struct")
	}

	t := v.Type()
	output := reflect.New(t).Elem()
	fieldSet := make(map[string]struct{})

	for _, field := range fieldsToIgnore {
		fieldSet[field] = struct{}{}
	}

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name

		if _, found := fieldSet[fieldName]; !found {
			output.Field(i).Set(v.Field(i))
		}
	}

	return output.Interface().(T), nil
}
