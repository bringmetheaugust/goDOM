package dom

import "goDOM/internal/errors"

type Document struct {
	root Element
}

func (d Document) New(rootEl Element) Document {
	d.root = rootEl

	return d
}

func (d Document) GetElementById(id string) (Element, error) {
	res, err := d.findByAttribute("id", id, d.root)

	if err != nil {
		return Element{}, errors.NotFoundErr{}
	} else {
		return res[0], nil
	}
}

func (d Document) GetElementsByClassName(class string) ([]Element, error) {
	return d.findByAttribute("class", class, d.root)
}

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
		return nil, errors.NotFoundErr{}
	} else {
		return matches, nil
	}
}
