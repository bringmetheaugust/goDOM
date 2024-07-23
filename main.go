package main

import (
	"fmt"
	"goDOM/internal/dom"
	"goDOM/internal/parser"
	"os"
)

func Create(data []byte) dom.Document {
	root := parser.ParseMarkup(string(data))

	return dom.Document{}.New(root)
}

func main() {
	dev()
}

func dev() {
	data, err := os.ReadFile("./mocks/test.html")

	if err != nil {
		panic(err)
	}

	document := Create(data)

	res, err1 := document.GetElementById("\"ouu\"")
	ress, err2 := document.GetElementsByClassName("\"lol\"")

	fmt.Println(res, ress, err1, err2)
}
