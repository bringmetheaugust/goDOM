package main

import (
	"fmt"
	"goDOM/internal/dom"
	"goDOM/internal/parser"
	"os"
)

func create(data []byte) dom.Document {
	root := parser.ParseMarkup(string(data))

	return dom.CreateDocument(root)
}

func main() {
	data, err := os.ReadFile("./mocks/test.html")

	if err != nil {
		panic(err)
	}

	document := create(data)

	res, _ := document.GetElementById("\"ouu\"")
	ress, _ := document.GetElementsByClassName("\"lol\"")

	fmt.Println(res, ress)
}
