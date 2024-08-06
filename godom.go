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

// Parsing markup, create and return DOM tree and DOM API.
//
// Example:
//
// document, err := goDom.Create(bytes)
// if err != nil {return} // if markup is invalid
func Create(data []byte) (*Document, error) {
	root, err := parse(string(data))

	if err != nil {
		return nil, err
	}

	return createDocument(root), nil
}
