package goDom

type attributes map[string]string
type classList []string
type children []Element

// DOM element with fields and element's DOM API.
type Element struct {
	TagName                string
	TextContent            string     // empty string if no content
	Attributes             attributes // nil if no attributes
	Children               children   // nil if no children
	ClassName              string     // string with CSS classes. empty string if no classes
	ClassList              classList  // nil if no classes
	Id                     string     // id attribute
	ParentElement          *Element   // nil if no parent element
	NextElementSibling     *Element   // nil if no next element sibling
	PreviousElementSibling *Element   // nil if no previous element sibling
	domSearchAPI
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

// Find element by query selector inside element. Exactly as yourElement.querySelector() in browser DOM.
//
// Examples:
//
//	element, err := yourElement.QuerySelector("div#lal .lol")
//	if err != nil {return} // if element doesn't exist inside element
//	fmt.Println(element) // print finded element
func (e Element) QuerySelector(queryStr string) (Element, error) {
	return e._querySelector(queryStr, e)
}

// Find elements by query selector inside element. Exactly as yourElement.querySelectorAll() in browser DOM.
//
// Examples:
//
//	elements, err := yourElement.QuerySelector("#my_lol .lolipop")
//	if err != nil {return} // if elements don't exist inside element
//	fmt.Println(elements) // print finded elements
func (e Element) QuerySelectorAll(queryStr string) ([]Element, error) {
	return e._querySelectorAll(queryStr, e)
}

// Find element by id inside element. Exactly as yourElement.getElementById() in browser DOM.
//
// Examples:
//
//	element, err := yourElement.GetElementById("piu")
//	if err != nil {return} // if element doesn't exist inside element
//	fmt.Println(element) // print finded element
func (e Element) GetElementById(id string) (Element, error) {
	return e._getElementById(id, e)
}

// Find elements by CSS class name inside element. Exactly as yourElement.getElementsByClassName() in browser DOM.
//
// Examples:
//
//	elements, err := yourElement.GetElementsByClassName(".lolipop")
//	if err != nil {return} // if elements don't exist inside element
//	fmt.Println(elements) // print finded elements
func (e Element) GetElementsByClassName(class string) ([]Element, error) {
	return e._getElementsByClassName(class, e)
}

// Find elements by tag name inside element. Exactly as yourElement.getElementsByTagName() in browser DOM.
//
// Examples:
//
//	elements, err := yourElement.GetElementsByTagName("li")
//	if err != nil {return} // if elements don't exist inside element
//	fmt.Println(elements) // print finded elements
func (e Element) GetElementsByTagName(tag string) ([]Element, error) {
	return e._getElementsByTagName(tag, e)
}
