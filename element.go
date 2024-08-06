package goDom

type attributes map[string]string

// DOM element with fields and element's API.
type Element struct {
	TagName       string
	TextContent   string // only own text/content
	Attributes    attributes
	Children      []Element
	ClassName     string
	ClassList     []string
	Id            string
	ParentElement *Element
}

// Get HTML attribute.
//
// Example:
//
// el, _ := document.GetElementById("lol")
// hrefAttr, err := el.GetAttribute("href")
// if err != nil {return} // attribute doesn't exist
// fmt.Println(hrefAttr) // print existed attribute value
func (e Element) GetAttribute(attr string) (string, error) {
	if val, ok := e.Attributes[attr]; ok {
		return val, nil
	}

	return "", notFoundErr{}
}

// Element has HTML attribute.
//
// Example:
//
// el, _ := document.GetElementById("lol")
// hasHrefAttr := el.HasAttribute("href")
// fmt.Println(hasHrefAttr) // print true if attribute existed
func (e Element) HasAttribute(attr string) bool {
	_, err := e.GetAttribute(attr)

	return err == nil
}
