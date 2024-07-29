package main

import (
	"goDOM/internal/dom"
	"goDOM/internal/parser"
)

func Create(data []byte) dom.Document {
	root, _ := parser.ParseHTML(string(data))

	return dom.CreateDocument(root)
}

func main() {
	dev()
}
