package goDom

import (
	"slices"
)

// DOM API search methods.
type domSearchAPI struct{}

func (api domSearchAPI) querySelector(queryStr string, el *Element) (*Element, error) {
	pq, err := parseQuery(queryStr)

	if err != nil {
		return nil, err
	}

	// Find one (first) element which matches all levels query.
	var findElementByQuery func(query, *Element) (*Element, error)
	findElementByQuery = func(q query, el *Element) (*Element, error) {
		var match *Element
		conditionFn := func(element *Element) bool {
			return elementMatchesQuery(q, element)
		}
		res, err := findOneByCondition(conditionFn, el)

		if err != nil {
			return nil, err
		}

		if q.child == nil {
			return res, nil
		}

		for _, el := range res.Children {
			e, err := findElementByQuery(*q.child, el)

			if err != nil {
				continue
			}

			match = e
		}

		if match == nil {
			return nil, notFoundErr{Params: queryStr}
		}

		return match, nil
	}

	return findElementByQuery(*pq, el)
}

func (api domSearchAPI) querySelectorAll(queryStr string, el *Element) ([]*Element, error) {
	queries, err := parseQueries(queryStr)

	if err != nil {
		return nil, err
	}

	var matches []*Element

	// Find elements which matches all levels query.
	var findElementsByQuery func(query, *Element) ([]*Element, error)
	findElementsByQuery = func(q query, el *Element) ([]*Element, error) {
		var matches []*Element

		// find elements which match first query level
		conditionFn := func(element *Element) bool {
			return elementMatchesQuery(q, element)
		}
		res, err := findAllByCondition(conditionFn, el)

		if err != nil {
			return nil, err
		}

		if q.child == nil {
			return res, nil
		}

		for _, match := range res {
			for _, e := range match.Children {
				elems, err := findElementsByQuery(*q.child, e)

				if err != nil {
					continue
				}

				matches = append(matches, elems...)
			}
		}

		return matches, nil
	}

	// paired queries
	for _, q := range queries {
		res, err := findElementsByQuery(q, el)

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

func (api domSearchAPI) getElementById(id string, el *Element) (*Element, error) {
	conditionFn := func(element *Element) bool {
		return element.Id == id
	}

	return findOneByCondition(conditionFn, el)
}

func (api domSearchAPI) getElementsByClassName(class string, el *Element) ([]*Element, error) {
	conditionFn := func(el *Element) bool {
		return slices.Contains(el.ClassList, class)
	}

	return findAllByCondition(conditionFn, el)
}

func (api domSearchAPI) getElementsByTagName(tag string, el *Element) ([]*Element, error) {
	conditionFn := func(element *Element) bool {
		return element.TagName == tag
	}

	return findAllByCondition(conditionFn, el)
}

func (api domSearchAPI) contains(root *Element, nested *Element) bool {
	conditionFn := func(e *Element) bool {
		return e == nested
	}

	_, err := findOneByCondition(conditionFn, root)

	return err == nil
}

// Check if element matches one level query.
func elementMatchesQuery(q query, el *Element) bool {
	if o := q.operator; o != "" {
		switch o {
		case query_operator_all:
			return true
		}
	} else {
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
	}

	return true
}

// Find elements by conditions.
func findAllByCondition(conditionFn func(*Element) bool, el *Element) ([]*Element, error) {
	var matches []*Element

	if conditionFn(el) {
		matches = append(matches, el)
	}

	for _, child := range el.Children {
		res, err := findAllByCondition(conditionFn, child)

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
func findOneByCondition(conditionFn func(*Element) bool, el *Element) (*Element, error) {
	if conditionFn(el) {
		return el, nil
	}

	for _, child := range el.Children {
		res, err := findOneByCondition(conditionFn, child)

		if err != nil {
			continue
		}

		return res, nil
	}

	return nil, notFoundErr{}
}
