package parser

import (
	"goDOM/internal/dom"
	"goDOM/internal/errors"
)

// Prepare and parse markup. Get DOM-like element tree.
func Parse(markup string) (*dom.Element, error) {
	markup = normalize(markup)

	if len(markup) == 0 {
		return &dom.Element{}, errors.InvalidRequest{Place: "markup is an empty string."}
	}

	parsedMarkup := parse(markup)

	return parsedMarkup, nil
}
