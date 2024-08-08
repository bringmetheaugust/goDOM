package goDom

// DOM tree with DOM API.
type Document struct {
	root Element
	domSearchAPI
}

// Find element by query selector. Exactly as document.querySelector() in browser DOM.
//
// Examples:
//
//	element, err := document.QuerySelector("div#lal .lol")
//	if err != nil {return} // if element doesn't exist in DOM tree
//	fmt.Println(element) // print finded element
func (d Document) QuerySelector(queryStr string) (Element, error) {
	return d._querySelector(queryStr, d.root)
}

// Find elements by query selector. Exactly as document.querySelectorAll() in browser DOM.
//
// Examples:
//
//	elements, err := document.QuerySelector("#my_lol .lolipop")
//	if err != nil {return} // if elements don't exist in DOM tree
//	fmt.Println(elements) // print finded elements
func (d Document) QuerySelectorAll(queryStr string) ([]Element, error) {
	return d._querySelectorAll(queryStr, d.root)
}

// Find element by id. Exactly as document.getElementById() in browser DOM.
//
// Examples:
//
//	element, err := document.GetElementById("piu")
//	if err != nil {return} // if element doesn't exist in DOM tree
//	fmt.Println(element) // print finded element
func (d Document) GetElementById(id string) (Element, error) {
	return d._getElementById(id, d.root)
}

// Find elements by CSS class name. Exactly as document.getElementsByClassName() in browser DOM.
//
// Examples:
//
//	elements, err := document.GetElementsByClassName(".lolipop")
//	if err != nil {return} // if elements don't exist in DOM tree
//	fmt.Println(elements) // print finded elements
func (d Document) GetElementsByClassName(class string) ([]Element, error) {
	return d._getElementsByClassName(class, d.root)
}

// Find elements by tag name. Exactly as document.getElementsByTagName() in browser DOM.
//
// Examples:
//
//	elements, err := document.GetElementsByTagName("li")
//	if err != nil {return} // if elements don't exist in DOM tree
//	fmt.Println(elements) // print finded elements
func (d Document) GetElementsByTagName(tag string) ([]Element, error) {
	return d._getElementsByTagName(tag, d.root)
}

// Create new document.
func createDocument(rootEl *Element) *Document {
	return &Document{root: *rootEl}
}
