package dom

import (
	"errors"
	"goDOM/internal/vdom"
)

func (d Document) GetElementById(id string) (vdom.Element, error) {
	res, err := d.findByAttribute("id", id, d.root)

	if err != nil {
		return vdom.Element{}, errors.New("404")
	} else {
		return res[0], nil
	}
}

func (d Document) GetElementsByClassName(class string) ([]vdom.Element, error) {
	return d.findByAttribute("class", class, d.root)
}
