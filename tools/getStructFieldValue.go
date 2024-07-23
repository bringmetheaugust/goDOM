package tools

import (
	"fmt"
	"reflect"
)

// Get dymanicly struct field.
func GetFieldValue(s interface{}, fieldName string) (interface{}, error) {
	v := reflect.ValueOf(s)

	// Проверяем, что переданный интерфейс является указателем на структуру
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("необходимо передать указатель на структуру")
	}

	// Получаем значение структуры
	v = v.Elem()

	// Получаем значение поля
	fieldVal := v.FieldByName(fieldName)

	if !fieldVal.IsValid() {
		return nil, fmt.Errorf("field %s doesnt exist", fieldName)
	}

	return fieldVal.Interface(), nil
}
