package goDom

import (
	"github.com/bringmetheaugust/goDOM/internal/dom"
	"github.com/bringmetheaugust/goDOM/internal/parser"
)

type Element dom.Element
type Document dom.Document

// Parsing markup, create and return DOM with element tree and API.
func Create(data []byte) (*dom.Document, error) {
	root, err := parser.Parse(string(data))

	if err != nil {
		return nil, err
	}

	return dom.CreateDocument(root), nil
}
