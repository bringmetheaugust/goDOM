package goDom

type attributes map[string]string
type classList []string
type children []Element

// DOM element with fields and element's API.
type Element struct {
	TagName       string
	TextContent   string // only own text/content
	Attributes    attributes
	Children      children
	ClassName     string
	ClassList     classList
	Id            string
	ParentElement *Element
}

// Get HTML attribute.
//
// Examples:
//
//	el, _ := document.GetElementById("lol")
//	hrefAttr, err := el.GetAttribute("href")
//	if err != nil {return} // attribute doesn't exist
//	fmt.Println(hrefAttr) // print existed attribute value
func (e Element) GetAttribute(attr string) (string, error) {
	if val, ok := e.Attributes[attr]; ok {
		return val, nil
	}

	return "", notFoundErr{Params: attr, Msg: "Attribute not found."}
}

// Element has HTML attribute.
//
// Examples:
//
//	el, _ := document.GetElementById("lol")
//	hasHrefAttr := el.HasAttribute("href")
//	fmt.Println(hasHrefAttr) // print true if attribute existed
func (e Element) HasAttribute(attr string) bool {
	_, err := e.GetAttribute(attr)

	return err == nil
}
