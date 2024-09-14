package goDom

import "slices"

type domAPIUtils struct{}

// check if element matches one level query.
func (u domAPIUtils) elementMatchesQuery(q query, el *Element) bool {
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

// find first element by conditions.
func (u domAPIUtils) findOneByCondition(conditionFn func(*Element) bool, el *Element) (*Element, error) {
	for _, child := range el.Children {
		if conditionFn(child) {
			return child, nil
		}

		res, err := u.findOneByCondition(conditionFn, child)

		if err != nil {
			continue
		}

		return res, nil
	}

	return nil, notFoundErr{}
}

// find all elements by conditions.
func (u domAPIUtils) findAllByCondition(conditionFn func(*Element) bool, el *Element) ([]*Element, error) {
	var matches []*Element

	for _, child := range el.Children {
		if conditionFn(child) {
			matches = append(matches, child)
		}

		res, err := u.findAllByCondition(conditionFn, child)

		if err == nil {
			matches = append(matches, res...)
		}
	}

	if len(matches) == 0 {
		return nil, notFoundErr{}
	}

	return matches, nil
}
