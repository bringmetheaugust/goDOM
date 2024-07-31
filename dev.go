package main

import (
	"fmt"
	"os"
)

func dev() {
	// data, _ := os.ReadFile("./test/query_selector.html")
	data, _ := os.ReadFile("./test/parse_html.html")

	document, _ := Create(data)

	fmt.Println(document)

	res1, err1 := document.GetElementById("ouu")
	res2, err2 := document.GetElementsByClassName("lol")
	res3, err3 := document.GetElementsByTagName("ul")
	res4, err4 := document.QuerySelector("#ou.lul .two li span")
	res5, err5 := document.QuerySelectorAll("li")

	fmt.Println(res1, err1)
	fmt.Println(res2, err2)
	fmt.Println(res3, err3)
	fmt.Println(res4, err4)
	fmt.Println(res5, err5)
}
