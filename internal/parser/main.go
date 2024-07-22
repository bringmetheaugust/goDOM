package parser

import "goDOM/internal/vdom"

// Parsing HTML markup.
// Return browser DOM-like struct.
func ParseMarkup(markup string) vdom.Element {
	element := parseElement(markup)

	return element
}
