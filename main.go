package main

import (
	"fmt"
	"goDOM/internal/dom"
	"goDOM/internal/parser"
	"os"
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

func dev() {
	data, _ := os.ReadFile("./mocks/test.html")

	document := Create(data)

	res1, err1 := document.GetElementById("ouu")
	res2, err2 := document.GetElementsByClassName("lol")
	res3, err3 := document.GetElementsByTagName("ul")
	// res4, err4 := document.QuerySelectorAll("div#lol.po")
	// res5, err5 := document.QuerySelectorAll(".po")
	// res6, err6 := document.QuerySelectorAll("div.piu.pou")
	// res7, err7 := document.QuerySelectorAll("span a#my_a.flex div.kill")

	fmt.Println(res1, err1)
	fmt.Println(res2, err2)
	fmt.Println(res3, err3)
	// fmt.Println(res4, err4)
	// fmt.Println(res5, err5)
	// fmt.Println(res6, err6)
	// fmt.Println(res7, err7)
}
