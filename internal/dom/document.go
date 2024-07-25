package dom

import (
	"goDOM/internal/errors"
	"goDOM/tools"
	"slices"
)

type Document struct {
	root Element
}

func (d Document) QuerySelector(queryStr string) (Element, error) {
	conditionFn, err := d.elementMatchesQuery(queryStr)

	if err != nil {
		return Element{}, err
	}

	return d.findOneByCondition(conditionFn, d.root)
}

func (d Document) QuerySelectorAll(queryStr string) ([]Element, error) {
	conditionFn, err := d.elementMatchesQuery(queryStr)

	if err != nil {
		return nil, err
	}

	return d.findAllByCondition(conditionFn, d.root)
}

func (d Document) GetElementById(id string) (Element, error) {
	res, err := d.findByField("Id", id, d.root)

	if err != nil {
		return Element{}, err
	}

	return res[0], nil
}

func (d Document) GetElementsByClassName(class string) ([]Element, error) {
	conditionFn := func(el Element) bool {
		return slices.Contains(el.ClassList, class)
	}

	return d.findAllByCondition(conditionFn, d.root)
}

func (d Document) GetElementsByTagName(tag string) ([]Element, error) {
	return d.findByField("TagName", tag, d.root)
}

// * Helpers

// TODO: deep query search
// Check if element matches query.
func (d Document) elementMatchesQuery(queryStr string) (func(Element) bool, error) {
	q, err := parseQuery(queryStr)

	conditionFn := func(el Element) bool {
		if q.tagName != "" && el.TagName != q.tagName {
			return false
		}

		if q.id != "" && el.Id != q.id {
			return false
		}

		// check if all classes from query contains element
		for _, class := range q.classList {
			if !slices.Contains(el.ClassList, class) {
				return false
			}
		}

		return true
	}

	return conditionFn, err
}

// Get elements by field. Only for cases when field's value is string.
func (d Document) findByField(field string, val string, el Element) ([]Element, error) {
	conditionFn := func(el Element) bool {
		fieldValue, err := tools.GetFieldValue(&el, field)

		if err != nil {
			return false
		}

		return fieldValue == val
	}

	return d.findAllByCondition(conditionFn, el)
}

// Get elements by attribute.
// func (d Document) findByAttribute(attr string, val string, el Element) ([]Element, error) {
// 	conditionFn := func(el Element) bool {
// 		for _, attrSearch := range el.Attributes {
// 			if attrSearch.Name == attr && attrSearch.Value == val {
// 				return true
// 			}
// 		}

// 		return false
// 	}

// 	return d.findAllByCondition(conditionFn, el)
// }

// Get results by conditions.
func (d Document) findAllByCondition(conditionFn func(Element) bool, el Element) ([]Element, error) {
	var matches []Element

	// If element satisfies the condition
	if conditionFn(el) {
		matches = append(matches, el)
	}

	// do the same for childrens (recursion)
	for _, child := range el.Children {
		res, err := d.findAllByCondition(conditionFn, child)

		if err != nil {
			continue
		}

		matches = append(matches, res...)
	}

	if len(matches) == 0 {
		return nil, errors.NotFound{}
	}

	return matches, nil
}

// Get first result by conditions.
func (d Document) findOneByCondition(conditionFn func(Element) bool, el Element) (Element, error) {
	// If element satisfies the condition
	if conditionFn(el) {
		return el, nil
	}

	// check childrens (recursion)
	for _, child := range el.Children {
		res, err := d.findOneByCondition(conditionFn, child)

		if err != nil {
			continue
		}

		return res, nil
	}

	return Element{}, errors.NotFound{}
}
