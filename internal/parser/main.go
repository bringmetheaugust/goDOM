package parser

import "goDOM/internal/dom"

// Parsing HTML markup.
// Return DOM-like struct.
func ParseMarkup(markup string) dom.Element {
	element := parseElement(markup, dom.Element{})

	return element
}
