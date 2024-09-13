// Parse HTML and get browser-like DOM and DOM API.
// Created only for reading DOM and getting information about elements.
// Doesn't have mathods to mutate DOM.
//
// # Parsing
//
// Uses net/html package to get HTML tokens.
//
// # DOM
// Create browser-like DOM with popular element's fields and methods to get inforamtion about selected element.
//
// # DOM API
// Created DOM has API to search elements by popular methods, like in browser DOM API.
//
// # jQuery
// Create jQuery-like API with popular methods and fields to get inforamtion about selected element.

package goDom

import (
	"strings"

	"golang.org/x/net/html"
)

// Parsing markup, create and return DOM tree with DOM API or jQuery.
// Return error if markup is invalid.
//
// Examples:
//
//	// getting DOM with DOM API
//	document, _, err := goDom.Create(bytes)
//	if err != nil {	// if markup is invalid
//
//	// getting jQuery
//	_, jQ, err := goDom.Create(bytes)
//	if err != nil {	// if markup is invalid
func Create(data []byte) (*Document, JQuery, error) {
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

	d, err := buildDOM(ch)

	if err != nil {
		return nil, nil, err
	}

	return d, createJQuery(d), nil
}
