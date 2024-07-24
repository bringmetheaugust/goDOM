package dom

import "goDOM/internal/errors"

type Element struct {
	TagName       string
	TextContent   string // only own text/content
	Attributes    []Attribute
	Children      []Element
	ClassName     string
	ClassList     []string
	FirstChild    *Element
	LastChild     *Element
	Id            string
	ParentElement *Element
}

// Get HTML attribute.
func (e Element) GetAttribute(attr string) (string, error) {
	for _, elAttr := range e.Attributes {
		if elAttr.Name == attr {
			return elAttr.Value, nil
		}
	}

	return "", &errors.NotFound{}
}

// Element has HTML attribute.
func (e Element) HasAttribute(attr string) bool {
	_, err := e.GetAttribute(attr)

	return err == nil
}
