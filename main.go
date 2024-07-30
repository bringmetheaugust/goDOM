package main

import (
	"goDOM/internal/dom"
	"goDOM/internal/parser"
)

// Parsing markup, create and return DOM with element tree and API.
func Create(data []byte) (*dom.Document, error) {
	root, err := parser.Parse(string(data))

	if err != nil {
		return nil, err
	}

	return dom.CreateDocument(root), nil
}

func main() {
	dev()
}
