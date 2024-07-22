package dom

import (
	"errors"
	"goDOM/internal/vdom"
)

func (d Document) findByAttribute(attr string, v string, el vdom.Element) ([]vdom.Element, error) {
	var matches []vdom.Element

	for _, attrSearch := range el.Attributes {
		if attrSearch.Name == attr && attrSearch.Value == v {
			matches = append(matches, el)
		}
	}

	for _, children := range el.Childrens {
		res, err := d.findByAttribute(attr, v, children)

		if err != nil {
			continue
		}

		matches = append(matches, res...)
	}

	if len(matches) == 0 {
		return nil, errors.New("404")
	} else {
		return matches, nil
	}
}
