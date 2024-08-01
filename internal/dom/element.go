package dom

import "github.com/bringmetheaugust/goDOM/internal/errors"

type Attributes map[string]string
type Element struct {
	TagName       string
	TextContent   string // only own text/content
	Attributes    Attributes
	Children      []Element
	ClassName     string
	ClassList     []string
	Id            string
	ParentElement *Element
}

// Get HTML attribute.
func (e Element) GetAttribute(attr string) (string, error) {
	if val, ok := e.Attributes[attr]; ok {
		return val, nil
	}

	return "", errors.NotFound{}
}

// Element has HTML attribute.
func (e Element) HasAttribute(attr string) bool {
	_, err := e.GetAttribute(attr)

	return err == nil
}
