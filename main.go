package main

import (
	"goDOM/internal/dom"
	"goDOM/internal/parser"
)

func Create(data []byte) (*dom.Document, error) {
	root, err := parser.ParseHTML(string(data))

	if err != nil {
		return nil, err
	}

	return dom.CreateDocument(root), nil
}

func main() {
	dev()
}
