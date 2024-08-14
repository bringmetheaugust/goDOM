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

// Check if tag is self-closing
func isSelfClosingTag(tag string) bool {
	return slices.Contains(selfClosingTags, tag)
}

// Parse markup. Get DOM-like element tree.
// Uses as downstream.
func parseMarkup(upStream chan token) *Element {
	var parentStack []Element
	var root Element
	var currEl *Element
	docType := html5

	for token := range upStream {
		switch {
		case token._type == token_content: // Element inner context
			if currEl != nil {
				currEl.TextContent += token.data

				continue
			}

			parentStack[len(parentStack)-1].TextContent += token.data
		case strings.HasPrefix(token.data, "</") && token._type == token_element: // tag is closing
			if currEl != nil {
				parentStack = append(parentStack, *currEl)
			}

			tagName := token.data[2 : len(token.data)-1]

			if len(parentStack) == 0 {
				panic("Error during parsing markup: unmatched closing tag. Please, report a bug." + token.data)
			}

			topFromParentStack := &parentStack[len(parentStack)-1]
			parentStack = parentStack[:len(parentStack)-1]

			if topFromParentStack.TagName != tagName {
				panic("Error during parsing markup: mismatched closing tag. Please, report a bug." + token.data)
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
		case strings.HasPrefix(token.data, "<!") && token._type == token_element: // doctype and comments
			if tokenLower := strings.ToLower(token.data); strings.HasPrefix(tokenLower, "<!doctype") {
				switch {
				case strings.HasPrefix(tokenLower, "<!doctype html public '-//w3c//dtd xhtml"):
					docType = xhtml
				default:
					docType = html5
				}
			}
		case strings.HasPrefix(token.data, "<") && token._type == token_element: // new element
			tag := parseTag(token.data)
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

			if strings.HasSuffix(token.data, "/>") || (docType == html5 && isSelfClosingTag(tag.name)) { // self-closing tags
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
		default: // Element inner context
			if currEl != nil {
				currEl.TextContent += token.data

				continue
			}

			parentStack[len(parentStack)-1].TextContent += token.data
		}
	}

	if len(parentStack) != 0 {
		panic("unmatched opening tags. Please, report a bug.")
	}

	return &root
}
