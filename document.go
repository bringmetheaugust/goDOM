package goDom

type docType string

const (
	html5 docType = "HTML5"
	xhtml docType = "XHTML"
)

// DOM tree with DOM API.
type Document struct {
	Title   *string    // https://developer.mozilla.org/en-US/docs/Web/API/Document/title
	Body    *Element   // https://developer.mozilla.org/en-US/docs/Web/API/Document/body
	Head    *Element   // https://developer.mozilla.org/en-US/docs/Web/API/Document/head
	Links   []*Element // https://developer.mozilla.org/en-US/docs/Web/API/Document/links
	Images  []*Element // https://developer.mozilla.org/en-US/docs/Web/API/Document/images
	Doctype docType    // https://developer.mozilla.org/en-US/docs/Web/API/Document/doctype
	root    *Element
	domSearchAPI
}

// Find element by query selector. Exactly as document.querySelector() in browser DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
//
// Examples:
//
//	element, err := document.QuerySelector("div#lal .lol")
//	if err != nil {return} // if element doesn't exist in DOM tree
//	fmt.Println(element) // print finded element
func (d *Document) QuerySelector(queryStr string) (*Element, error) {
	return d.querySelector(queryStr, d.root)
}

// Find elements by query selector. Exactly as document.querySelectorAll() in browser DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll
//
// Examples:
//
//	elements, err := document.QuerySelector("#my_lol .lolipop")
//	if err != nil {return} // if elements don't exist in DOM tree
//	fmt.Println(elements) // print finded elements
func (d *Document) QuerySelectorAll(queryStr string) ([]*Element, error) {
	return d.querySelectorAll(queryStr, d.root)
}

// Find element by id. Exactly as document.getElementById() in browser DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
//
// Examples:
//
//	element, err := document.GetElementById("piu")
//	if err != nil {return} // if element doesn't exist in DOM tree
//	fmt.Println(element) // print finded element
func (d *Document) GetElementById(id string) (*Element, error) {
	return d.getElementById(id, d.root)
}

// Find elements by CSS class name. Exactly as document.getElementsByClassName() in browser DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName
//
// Examples:
//
//	elements, err := document.GetElementsByClassName(".lolipop")
//	if err != nil {return} // if elements don't exist in DOM tree
//	fmt.Println(elements) // print finded elements
func (d *Document) GetElementsByClassName(class string) ([]*Element, error) {
	return d.getElementsByClassName(class, d.root)
}

// Find elements by tag name. Exactly as document.getElementsByTagName() in browser DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName
//
// Examples:
//
//	elements, err := document.GetElementsByTagName("li")
//	if err != nil {return} // if elements don't exist in DOM tree
//	fmt.Println(elements) // print finded elements
func (d *Document) GetElementsByTagName(tag string) ([]*Element, error) {
	return d.getElementsByTagName(tag, d.root)
}
