// Parse HTML and get browser-like DOM and DOM API.
// Created only for reading DOM and getting information about elements.
// Doesn't have mathods to mutate DOM.
//
// # DOM
// Create browser-like DOM with popular element's fields and methods to get inforamtion about selected element.
//
// # DOM API
// Created DOM has API to search elements by popular methods, like in browser DOM API.

package goDom

import (
	"strings"

	"golang.org/x/net/html"
)

// Parsing markup, create and return DOM tree with DOM API.
//
// Examples:
//
//	document, err := goDom.Create(bytes)
func Create(data []byte) (*Document, error) {
	ch := make(chan html.Token)
	t := html.NewTokenizer(strings.NewReader(string(data)))

	go func() {
		for {
			tt := t.Next()

			if tt == html.ErrorToken {
				close(ch)
				break
			}

			ch <- t.Token()
		}
	}()

	root, err := buildDOM(ch)

	if err != nil {
		return nil, err
	}

	return createDocument(root), nil
}
