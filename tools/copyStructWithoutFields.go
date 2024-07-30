package tools

import "reflect"

func CopyStructWithoutFields(input interface{}, fieldsToIgnore []string) interface{} {
	v := reflect.ValueOf(input)
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

	return output.Interface()
}
