// Provide method to parse HTML and get browser-like DOM and DOM API.
// Created only for reading DOM and getting information about elements.
// Doesn't have mathods to mutate DOM.
//
// # Parsing HTML
//
// Parser get DOM tree of elements. It's simple mechanism without creating nodes.
//
// # DOM
// Create browser-like DOM with popular element's fields and methods to get inforamtion about selected element.
//
// # DOM API
// Created DOM has API to search elements by popular methods, like in browser DOM API.

package goDom

// Prepare and parsing markup, create and return DOM tree and DOM API.
//
// Examples:
//
//	document, err := goDom.Create(bytes)
//	if err != nil {return} // if markup is invalid
func Create(data []byte) (*Document, error) {
	chMarkupLine := make(chan string)
	chTokens := make(chan string)

	go normalize(string(data), chMarkupLine)
	go tokenize(chMarkupLine, chTokens)
	root := parseMarkup(chTokens)

	return createDocument(root), nil
}
