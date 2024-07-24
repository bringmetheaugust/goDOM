package dom

import (
	"goDOM/internal/errors"
	"goDOM/tools"
)

type Document struct {
	root Element
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
func (d Document) findByField(field string, val string, el Element) ([]Element, error) {
	conditionFn := func(el Element) bool {
		fieldValue, err := tools.GetFieldValue(&el, field)

		if err != nil {
			return false
		}

		return fieldValue == val
	}

	return d.findByCondition(conditionFn, el)
}

// Get element by attribute.
func (d Document) findByAttribute(attr string, val string, el Element) ([]Element, error) {
	conditionFn := func(el Element) bool {
		for _, attrSearch := range el.Attributes {
			if attrSearch.Name == attr && attrSearch.Value == val {
				return true
			}
		}

		return false
	}

	return d.findByCondition(conditionFn, el)
}

// Get result by conditions, accamulate result.
func (d Document) findByCondition(conditionFn func(Element) bool, el Element) ([]Element, error) {
	var matches []Element

	// If element satisfies the condition
	if conditionFn(el) {
		matches = append(matches, el)
	}

	// do the same for childrens (recursion)
	for _, child := range el.Children {
		res, err := d.findByCondition(conditionFn, child)

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
