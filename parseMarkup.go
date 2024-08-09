package goDom

import (
	"slices"
	"strings"
)

type docType string

const (
	html5 docType = "HTML5"
	xhtml docType = "XHTML"
)

// self closing tags in HTML5
var selfClosingTags = []string{
	"area", "base", "br", "col", "embed", "hr", "img", "input",
	"link", "meta", "param", "source", "track", "wbr",
}

// Prepare and parse markup. Get DOM-like element tree.
func parse(markup string) (*Element, error) {
	markup = normalize(markup)

	if len(markup) == 0 {
		return &Element{}, invalidRequestErr{Place: "markup is an empty string."}
	}

	ch := make(chan string)

	go tokenize(markup, ch)
	parsedMarkup := parseMarkup(ch)

	return parsedMarkup, nil
}

// Check if tag is self-closing
func isSelfClosingTag(tag string) bool {
	return slices.Contains(selfClosingTags, tag)
}

// Parse markup. Get DOM-like element tree.
func parseMarkup(ch chan string) *Element {
	var parentStack []Element
	var root Element
	var currEl *Element
	docType := html5

	for token := range ch {
		switch {
		case strings.HasPrefix(token, "</"): // tag is closing
			if currEl != nil {
				parentStack = append(parentStack, *currEl)
			}

			tagName := token[2 : len(token)-1]

			if len(parentStack) == 0 {
				panic("Error during parsing markup: unmatched closing tag.")
			}

			topFromParentStack := &parentStack[len(parentStack)-1]
			parentStack = parentStack[:len(parentStack)-1]

			if topFromParentStack.TagName != tagName {
				panic("Error during parsing markup: mismatched closing tag.")
			}

			if currEl != nil {
				topFromParentStack.TextContent = currEl.TextContent
			}

			if len(parentStack) > 0 {
				parent := &parentStack[len(parentStack)-1]
				parent.Children = append(parent.Children, *topFromParentStack)
			} else {
				root = *topFromParentStack
			}

			currEl = nil

			continue
		case strings.HasPrefix(token, "<!"): // doctype and comments
			if tokenLower := strings.ToLower(token); strings.HasPrefix(tokenLower, "<!doctype") {
				switch {
				case strings.HasPrefix(tokenLower, "<!doctype html public '-//w3c//dtd xhtml"):
					docType = xhtml
				default:
					docType = html5
				}
			}

			continue
		case strings.HasPrefix(token, "<"): // new element
			tag := parseTag(token)
			newEl := Element{TagName: tag.name, Attributes: tag.attributes}

			for k, v := range tag.attributes {
				switch {
				case k == "class":
					newEl.ClassName = v
					newEl.ClassList = strings.Fields(v)

					continue
				case k == "id":
					newEl.Id = v

					continue
				}
			}

			if strings.HasSuffix(token, "/>") || (docType == html5 && isSelfClosingTag(tag.name)) { // self-closing tags
				if currEl != nil {
					currEl.Children = append(currEl.Children, newEl)
				} else {
					topFromParentStack := &parentStack[len(parentStack)-1]
					topFromParentStack.Children = append(topFromParentStack.Children, newEl)
				}
			} else { // tag opening ends
				if currEl != nil {
					parentStack = append(parentStack, *currEl)
					currEl.Children = append(currEl.Children, newEl)
					newEl.ParentElement = currEl
				}

				currEl = &newEl
			}

			continue
		default: // Element inner context
			if currEl != nil {
				currEl.TextContent += token

				continue
			}

			parentStack[len(parentStack)-1].TextContent += token
		}
	}

	if len(parentStack) != 0 {
		panic("unmatched opening tags")
	}

	return &root
}
