package goDom

import (
	"slices"

	"github.com/bringmetheaugust/goDOM/tools"
)

// DOM API search methods.
type domSearchAPI struct{}

func (a domSearchAPI) _querySelector(queryStr string, el Element) (Element, error) {
	q, err := parseQuery(queryStr)

	if err != nil {
		return Element{}, err
	}

	return a.findElementByQuery(*q, el)
}

func (a domSearchAPI) _querySelectorAll(queryStr string, el Element) ([]Element, error) {
	queries, err := parseQueries(queryStr)

	if err != nil {
		return nil, err
	}

	var matches []Element

	for _, q := range queries {
		res, err := a.findElementsByQuery(q, el)

		if err != nil {
			continue
		}

		matches = append(matches, res...)
	}

	if len(matches) == 0 {
		return nil, notFoundErr{Params: "query: " + queryStr}
	}

	return matches, nil
}

func (a domSearchAPI) _getElementById(id string, el Element) (Element, error) {
	res, err := a.findByField("Id", id, el)

	if err != nil {
		return Element{}, err
	}

	return res[0], nil
}

func (a domSearchAPI) _getElementsByClassName(class string, el Element) ([]Element, error) {
	conditionFn := func(el Element) bool {
		return slices.Contains(el.ClassList, class)
	}

	return a.findAllByCondition(conditionFn, el)
}

func (a domSearchAPI) _getElementsByTagName(tag string, el Element) ([]Element, error) {
	return a.findByField("TagName", tag, el)
}

// Check if element matches one level query.
func (a domSearchAPI) elementMatchesQuery(q query) func(Element) bool {
	conditionFn := func(el Element) bool {
		if q.tagName != "" && el.TagName != q.tagName {
			return false
		}

		if q.id != "" && el.Id != q.id {
			return false
		}

		// check if each class from query contains element
		for _, class := range q.classList {
			if !slices.Contains(el.ClassList, class) {
				return false
			}
		}

		// check if each attribute from query contains element
		for k, v := range q.attributes {
			attr, ok := el.Attributes[k]

			if !ok || (v != "" && v == attr) {
				return false
			}
		}

		return true
	}

	return conditionFn
}

// Find one (first) element which matches all levels query.
func (a domSearchAPI) findElementByQuery(q query, el Element) (Element, error) {
	conditionFn := a.elementMatchesQuery(q)
	res, err := a.findOneByCondition(conditionFn, el)

	if err != nil {
		return Element{}, err
	}

	if q.child == nil {
		return res, nil
	}

	return a.findElementByQuery(*q.child, res)
}

// Find elements which matches all levels query.
func (a domSearchAPI) findElementsByQuery(q query, el Element) ([]Element, error) {
	var matches []Element

	// find elements which match first query level
	conditionFn := a.elementMatchesQuery(q)
	res, err := a.findAllByCondition(conditionFn, el)

	if err != nil {
		return nil, err
	}

	if q.child == nil {
		return res, nil
	}

	for _, match := range res {
		res, err := a.findElementsByQuery(*q.child, match)

		if err != nil {
			continue
		}

		matches = append(matches, res...)
	}

	return matches, nil
}

// FInd elements by field. Only for cases when field's value is string.
func (a domSearchAPI) findByField(field string, val string, el Element) ([]Element, error) {
	conditionFn := func(el Element) bool {
		fieldValue, err := tools.GetFieldValue(&el, field)

		if err != nil {
			return false
		}

		return fieldValue == val
	}

	return a.findAllByCondition(conditionFn, el)
}

// Find elements by conditions.
func (a domSearchAPI) findAllByCondition(conditionFn func(Element) bool, el Element) ([]Element, error) {
	var matches []Element

	if conditionFn(el) {
		matches = append(matches, el)
	}

	for _, child := range el.Children {
		res, err := a.findAllByCondition(conditionFn, child)

		if err == nil {
			matches = append(matches, res...)
		}
	}

	if len(matches) == 0 {
		return nil, notFoundErr{}
	}

	return matches, nil
}

// Find first element by conditions.
// ! Can't make it async. Async will return any randomly gotten element because goroutines are сhaotic.
// ! This is not correct cause native DOM always searches element in the nearest subtree.
func (a domSearchAPI) findOneByCondition(conditionFn func(Element) bool, el Element) (Element, error) {
	if conditionFn(el) {
		return el, nil
	}

	for _, child := range el.Children {
		res, err := a.findOneByCondition(conditionFn, child)

		if err != nil {
			continue
		}

		return res, nil
	}

	return Element{}, notFoundErr{}
}

// Find elements by attribute.
// Temporary deprecated.
// func (a domSearchAPI) findByAttribute(attr string, val string, el Element) ([]Element, error) {
// 	conditionFn := func(el Element) bool {
// 		if v, ok := el.Attributes[attr]; ok && v == val {
// 			return true
// 		}

// 		return false
// 	}

// 	return a.findAllByCondition(conditionFn, el)
// }
