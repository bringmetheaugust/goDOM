package main

import (
	"fmt"
	"io"
	"net/http"

	goDom "github.com/bringmetheaugust/goDOM"
)

func main() {
	r, err := http.Get("https://stackoverflow.com")

	if err != nil || r.StatusCode >= 400 {
		fmt.Printf("Cann't get site: %d", err)
		return
	}

	defer r.Body.Close()

	bytes, _ := io.ReadAll(r.Body)
	document, _ := goDom.Create(bytes) // create document

	// find DOM elements as posts (.s-post-summary) and get their link elements (.s-post-summary--content-title a)
	posts, err := document.QuerySelectorAll(".s-post-summary .s-post-summary--content-title a")

	// ! Remember that sites can detect You as a bot and return unexpected HTML response
	if err != nil {
		fmt.Println("Ooops, posts not found.")
		return
	}

	// loop link elements (HTMLAnchorElement)
	for _, c := range posts {
		attr, err := c.GetAttribute("href") // get href attribute

		if err != nil {
			fmt.Println("Ooops, href not found.")
			return
		}

		fmt.Println(attr)
	}
}
