package main

import (
	"fmt"
	htmlparser "goDOM/internal/htmlParser"
	"os"
)

func main() {
	data, err := os.ReadFile("./mocks/test.html")

	if err != nil {
		panic(err)
	}

	fmt.Println(htmlparser.Parse(string(data)))
}
