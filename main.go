package main

import (
	"goDOM/internal/dom"
	"goDOM/internal/parser"
	"strings"
)

func Create(data []byte) dom.Document {
	unescapedStr := strings.ReplaceAll(string(data), `"`, `'`) // cause string() shielding /"
	root := parser.ParseMarkup(string(unescapedStr))

	return dom.CreateDocument(root)
}

func main() {
	dev()
}
