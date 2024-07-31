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
	q, err := parseQuery(queryStr)

	if err != nil {
		return Element{}, err
	}

	return d.findElementByQuery(*q, d.root)
}

func (d Document) QuerySelectorAll(queryStr string) ([]Element, error) {
	q, err := parseQuery(queryStr)

	if err != nil {
		return nil, err
	}

	res, err := d.findElementsByQuery(*q, d.root)

	if len(res) == 0 {
		return nil, err
	}

	return res, nil
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

// Check if element matches one level query.
func (d Document) elementMatchesQuery(q query) func(Element) bool {
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

	return conditionFn
}

// Find one (first) element which matches all levels query.
func (d Document) findElementByQuery(q query, el Element) (Element, error) {
	conditionFn := d.elementMatchesQuery(q)
	res, err := d.findOneByCondition(conditionFn, el)

	if err != nil {
		return Element{}, err
	}

	if q.child == nil {
		return res, nil
	}

	return d.findElementByQuery(*q.child, res)
}

// Find elements which matches all levels query.
func (d Document) findElementsByQuery(q query, el Element) ([]Element, error) {
	var matches []Element

	conditionFn := d.elementMatchesQuery(q)
	res, err := d.findAllByCondition(conditionFn, el)

	if err != nil {
		return nil, err
	}

	if q.child == nil {
		return res, nil
	}

	for _, match := range res {
		res, err := d.findElementsByQuery(*q.child, match)

		if err != nil {
			continue
		}

		matches = append(matches, res...)
	}

	return matches, nil
}

// FInd elements by field. Only for cases when field's value is string.
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

// Find elements by conditions.
func (d Document) findAllByCondition(conditionFn func(Element) bool, el Element) ([]Element, error) {
	var matches []Element

	if conditionFn(el) {
		matches = append(matches, el)
	}

	for _, child := range el.Children {
		res, err := d.findAllByCondition(conditionFn, child)

		if err == nil {
			matches = append(matches, res...)
		}
	}

	if len(matches) == 0 {
		return nil, errors.NotFound{}
	}

	return matches, nil
}

// Find first element by conditions.
func (d Document) findOneByCondition(conditionFn func(Element) bool, el Element) (Element, error) {
	if conditionFn(el) {
		return el, nil
	}

	for _, child := range el.Children {
		res, err := d.findOneByCondition(conditionFn, child)

		if err != nil {
			continue
		}

		return res, nil
	}

	return Element{}, errors.NotFound{}
}

// Find elements by attribute.
// Temporary deprecated.
// func (d Document) findByAttribute(attr string, val string, el Element) ([]Element, error) {
// 	conditionFn := func(el Element) bool {
// 		if v, ok := el.Attributes[attr]; ok && v == val {
// 			return true
// 		}

// 		return false
// 	}

// 	return d.findAllByCondition(conditionFn, el)
// }
