package main

import (
	"fmt"
	"goDOM/internal/dom"
	"goDOM/internal/parser"
	"os"
)

func Create(data []byte) dom.Document {
	root := parser.ParseMarkup(string(data))

	return dom.CreateDocument(root)
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

	res1, err1 := document.GetElementById("\"ouu\"")
	res2, err2 := document.GetElementsByClassName("\"lol")
	res3, err3 := document.GetElementsByTagName("ul")

	fmt.Println(res1, err1)
	fmt.Println(res2, err2)
	fmt.Println(res3, err3)
}
