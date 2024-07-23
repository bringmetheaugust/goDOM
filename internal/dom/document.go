package dom

import (
	"goDOM/internal/errors"
	"goDOM/tools"
)

type Document struct {
	root Element
}

// Create new document.
func (d Document) New(rootEl Element) Document {
	d.root = rootEl

	return d
}

func (d Document) GetElementById(id string) (Element, error) {
	res, err := d.findByField("Id", id, d.root)

	if err != nil {
		return Element{}, &errors.NotFound{}
	}

	return res[0], nil
}

// TODO doesn't work
func (d Document) GetElementsByClassName(class string) ([]Element, error) {
	return d.findByAttribute("class", class, d.root)
}

func (d Document) GetElementsByTagName(tag string) ([]Element, error) {
	return d.findByField("TagName", tag, d.root)
}

// Get element by field. Only for cases when field's value is string.
func (d Document) findByField(field string, v string, el Element) ([]Element, error) {
	var matches []Element

	fieldName, err := tools.GetFieldValue(&el, field)

	if err == nil && fieldName == v {
		matches = append(matches, el)
	}

	for _, children := range el.Children {
		res, err := d.findByField(field, v, children)

		if err != nil {
			continue
		}

		matches = append(matches, res...)
	}

	if len(matches) == 0 {
		return nil, &errors.NotFound{}
	}

	return matches, nil
}

// Get element by attribute.
func (d Document) findByAttribute(attr string, v string, el Element) ([]Element, error) {
	var matches []Element

	for _, attrSearch := range el.Attributes {
		if attrSearch.Name == attr && attrSearch.Value == v {
			matches = append(matches, el)
		}
	}

	for _, children := range el.Children {
		res, err := d.findByAttribute(attr, v, children)

		if err != nil {
			continue
		}

		matches = append(matches, res...)
	}

	if len(matches) == 0 {
		return nil, &errors.NotFound{}
	}

	return matches, nil
}

// TODO Optimization

// Обобщённая функция для поиска элементов
// func (d Document) findByCondition(condition func(Element) bool, el Element) ([]Element, error) {
// 	var matches []Element

// 	// Проверяем условие для текущего элемента
// 	if condition(el) {
// 		matches = append(matches, el)
// 	}

// 	// Рекурсивно проверяем детей
// 	for _, child := range el.Children {
// 		res, err := d.findByCondition(condition, child)
// 		if err != nil {
// 			continue
// 		}
// 		matches = append(matches, res...)
// 	}

// 	// Если нет совпадений, возвращаем ошибку
// 	if len(matches) == 0 {
// 		return nil, NotFoundErr
// 	}

// 	return matches, nil
// }

// // Функция для поиска по полю
// func (d Document) findByField(field string, value string, el Element) ([]Element, error) {
// 	condition := func(el Element) bool {
// 		fieldValue, err := getFieldValue(&el, field)
// 		if err != nil {
// 			return false
// 		}
// 		return fieldValue == value
// 	}
// 	return d.findByCondition(condition, el)
// }

// // Функция для поиска по атрибуту
// func (d Document) findByAttribute(attr string, value string, el Element) ([]Element, error) {
// 	condition := func(el Element) bool {
// 		for _, attrSearch := range el.Attributes {
// 			if attrSearch.Name == attr && attrSearch.Value == value {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	return d.findByCondition(condition, el)
// }
