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

// Get DOM-like element tree.
// Uses as downstream.
func buildDOM(upStream chan token) *Element {
	var parentStack []Element
	var root Element
	var currEl *Element
	docType := html5

	for token := range upStream {
		switch {
		case token._type == node_meta:
			if tokenLower := strings.ToLower(token.data); strings.HasPrefix(tokenLower, "<!doctype") {
				switch {
				case strings.HasPrefix(tokenLower, "<!doctype html public '-//w3c//dtd xhtml"):
					docType = xhtml
				default:
					docType = html5
				}
			}
		case token._type == node_text:
			if currEl != nil {
				currEl.TextContent += token.data

				continue
			}

			parentStack[len(parentStack)-1].TextContent += token.data
		case token._type == node_element && token.isClosing: // element closing
			if currEl != nil {
				parentStack = append(parentStack, *currEl)
			}

			if len(parentStack) == 0 {
				panic("Error during parsing markup: unmatched closing tag. Please, report a bug.")
			}

			// tagName := token.tag.name
			topFromParentStack := &parentStack[len(parentStack)-1]
			parentStack = parentStack[:len(parentStack)-1]

			// if topFromParentStack.TagName != tagName {
			// 	panic("Error during parsing markup: mismatched closing tag. Please, report a bug.")
			// }

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
		case token._type == node_element && !token.isClosing: // new element
			tag := token.tag
			newEl := Element{TagName: tag.name, Attributes: tag.attributes}

			for k, v := range tag.attributes {
				switch {
				case k == "class":
					newEl.ClassName = v
					newEl.ClassList = strings.Fields(v)
				case k == "id":
					newEl.Id = v
				}
			}

			if tag.selfClosing || (docType == html5 && isSelfClosingTag(tag.name)) { // self-closing tags
				if currEl != nil {
					currEl.Children = append(currEl.Children, newEl)
				} else {
					topFromParentStack := &parentStack[len(parentStack)-1]
					topFromParentStack.Children = append(topFromParentStack.Children, newEl)
				}
			} else { // tag ends
				if currEl != nil {
					parentStack = append(parentStack, *currEl)
					currEl.Children = append(currEl.Children, newEl)
					newEl.ParentElement = currEl
				}

				currEl = &newEl
			}
		}
	}

	if len(parentStack) != 0 {
		panic("unmatched opening tags. Please, report a bug.")
	}

	return &root
}
