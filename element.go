package goDom

type attributes map[string]string
type classList []string
type children []*Element

// DOM element with fields and element's DOM API.
type Element struct {
	TagName                string     // https://developer.mozilla.org/en-US/docs/Web/API/Element/tagName
	TextContent            string     // https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
	Attributes             attributes // https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
	Children               children   // https://developer.mozilla.org/en-US/docs/Web/API/Element/children
	ClassName              string     // https://developer.mozilla.org/en-US/docs/Web/API/Element/className
	ClassList              classList  // https://developer.mozilla.org/en-US/docs/Web/API/Element/classList
	Id                     string     // https://developer.mozilla.org/en-US/docs/Web/API/Element/id
	ParentElement          *Element   // https://developer.mozilla.org/en-US/docs/Web/API/Node/parentElement
	NextElementSibling     *Element   // https://developer.mozilla.org/en-US/docs/Web/API/Element/nextElementSibling
	PreviousElementSibling *Element   // https://developer.mozilla.org/en-US/docs/Web/API/Element/previousElementSibling
	domSearchAPI
}

// Get HTML attribute.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getAttribute
//
// Examples:
//
//	el, _ := document.GetElementById("lol")
//	hrefAttr, err := el.GetAttribute("href")
//	if err != nil {return} // attribute doesn't exist
//	fmt.Println(hrefAttr) // print existed attribute value
func (e *Element) GetAttribute(attr string) (string, error) {
	if val, ok := e.Attributes[attr]; ok {
		return val, nil
	}

	return "", notFoundErr{Params: attr, Msg: "Attribute not found."}
}

// Element has HTML attribute.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/hasAttribute
//
// Examples:
//
//	el, _ := document.GetElementById("lol")
//	hasHrefAttr := el.HasAttribute("href")
//	fmt.Println(hasHrefAttr) // print true if attribute existed
func (e *Element) HasAttribute(attr string) bool {
	_, err := e.GetAttribute(attr)

	return err == nil
}

// Find element by query selector inside element. Exactly as yourElement.querySelector() in browser DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
//
// Examples:
//
//	element, err := yourElement.QuerySelector("div#lal .lol")
//	if err != nil {return} // if element doesn't exist inside element
//	fmt.Println(element) // print finded element
func (e *Element) QuerySelector(queryStr string) (*Element, error) {
	return e.querySelector(queryStr, e)
}

// Find elements by query selector inside element. Exactly as yourElement.querySelectorAll() in browser DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll
//
// Examples:
//
//	elements, err := yourElement.QuerySelector("#my_lol .lolipop")
//	if err != nil {return} // if elements don't exist inside element
//	fmt.Println(elements) // print finded elements
func (e *Element) QuerySelectorAll(queryStr string) ([]*Element, error) {
	return e.querySelectorAll(queryStr, e)
}

// Find element by id inside element. Exactly as yourElement.getElementById() in browser DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
//
// Examples:
//
//	element, err := yourElement.GetElementById("piu")
//	if err != nil {return} // if element doesn't exist inside element
//	fmt.Println(element) // print finded element
func (e *Element) GetElementById(id string) (*Element, error) {
	return e.getElementById(id, e)
}

// Find elements by CSS class name inside element. Exactly as yourElement.getElementsByClassName() in browser DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName
//
// Examples:
//
//	elements, err := yourElement.GetElementsByClassName(".lolipop")
//	if err != nil {return} // if elements don't exist inside element
//	fmt.Println(elements) // print finded elements
func (e *Element) GetElementsByClassName(class string) ([]*Element, error) {
	return e.getElementsByClassName(class, e)
}

// Find elements by tag name inside element. Exactly as yourElement.getElementsByTagName() in browser DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName
//
// Examples:
//
//	elements, err := yourElement.GetElementsByTagName("li")
//	if err != nil {return} // if elements don't exist inside element
//	fmt.Println(elements) // print finded elements
func (e *Element) GetElementsByTagName(tag string) ([]*Element, error) {
	return e.getElementsByTagName(tag, e)
}

// If DOM element contains inside itself another DOM element.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/contains
//
// Examples:
//
//	parent, _ := document.QuerySelector(".pups")
//	child, _ := document.GetElementById("wee")
//	parent.Contains(child) // return true is parent contains inside itself child
func (e *Element) Contains(el *Element) bool {
	return e.contains(e, el)
}
