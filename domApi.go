package goDom

import (
	"slices"
)

// DOM API search methods.
type domAPI struct {
	domAPIUtils
}

// find one element by query selector
func (api domAPI) querySelector(queryStr string, el *Element) (*Element, error) {
	pq, err := parseQuery(queryStr)

	if err != nil {
		return nil, err
	}

	// Find one (first) element which matches all levels query.
	var findElementByQuery func(query, *Element) (*Element, error)
	findElementByQuery = func(q query, el *Element) (*Element, error) {
		var match *Element
		conditionFn := func(element *Element) bool {
			return api.elementMatchesQuery(q, element)
		}

		res, err := api.findOneByCondition(conditionFn, el)

		if err != nil {
			return nil, err
		}

		if q.child == nil {
			return res, nil
		}

		for _, el := range res.Children {
			if e, err := findElementByQuery(*q.child, el); err == nil {
				match = e
			}
		}

		if match == nil {
			return nil, notFoundErr{Params: queryStr}
		}

		return match, nil
	}

	return findElementByQuery(*pq, el)
}

// find elements by query selector
func (api domAPI) querySelectorAll(queryStr string, el *Element) ([]*Element, error) {
	queries, err := parseQueries(queryStr)

	if err != nil {
		return nil, err
	}

	// find elements which matches all levels query.
	var findElementsByQuery func(query, *Element) ([]*Element, error)
	findElementsByQuery = func(q query, el *Element) ([]*Element, error) {
		var matches []*Element

		// find elements which match first query level
		conditionFn := func(element *Element) bool {
			return api.elementMatchesQuery(q, element)
		}

		res, err := api.findAllByCondition(conditionFn, el)

		if err != nil {
			return nil, err
		}

		if q.child == nil {
			return res, nil
		}

		for _, e := range res {
			if elems, err := findElementsByQuery(*q.child, e); err == nil {
				matches = append(matches, elems...)
			}
		}

		return matches, nil
	}

	var matches []*Element

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

// get element by Id attribute
func (api domAPI) getElementById(id string, el *Element) (*Element, error) {
	conditionFn := func(element *Element) bool {
		return element.Id == id
	}

	return api.findOneByCondition(conditionFn, el)
}

// get elements by class name
func (api domAPI) getElementsByClassName(class string, el *Element) ([]*Element, error) {
	conditionFn := func(el *Element) bool {
		return slices.Contains(el.ClassList, class)
	}

	return api.findAllByCondition(conditionFn, el)
}

// get elements by tag name
func (api domAPI) getElementsByTagName(tag string, el *Element) ([]*Element, error) {
	conditionFn := func(element *Element) bool {
		return element.TagName == tag
	}

	return api.findAllByCondition(conditionFn, el)
}

// check if one element contains another element
func (api domAPI) contains(root *Element, nested *Element) bool {
	conditionFn := func(e *Element) bool {
		return e == nested
	}

	_, err := api.findOneByCondition(conditionFn, root)

	return err == nil
}
