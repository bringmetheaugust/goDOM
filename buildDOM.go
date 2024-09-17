package goDom

import (
	"slices"
	"strings"

	"golang.org/x/net/html"
)

// self closing tags in HTML5
var selfClosingTags = []string{"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr"}

// Check if tag is self-closing
func isSelfClosingTag(tag string) bool {
	return slices.Contains(selfClosingTags, tag)
}

// Get DOM-like element tree.
// Uses as downstream.
func buildDOM(upStream chan html.Token) (*Document, error) {
	var doc Document
	var parentStack []*Element
	var currEl *Element

rLoop:
	for t := range upStream {
		switch t.Type {
		case html.CommentToken:
			continue rLoop
		case html.DoctypeToken:
			switch tLow := strings.ToLower(t.Data); {
			case strings.HasPrefix(tLow, "<!doctype html public '-//w3c//dtd xhtml"):
				doc.Doctype = xhtml
			default:
				doc.Doctype = html5
			}
		case html.TextToken:
			if str := strings.TrimSpace(t.Data); str != "" {
				if currEl != nil {
					currEl.TextContent += str
				} else {
					parentStack[len(parentStack)-1].TextContent += str
				}
			}
		case html.EndTagToken:
			if currEl != nil {
				parentStack = append(parentStack, currEl)
			}

			topFromParentStack := parentStack[len(parentStack)-1]
			parentStack = parentStack[:len(parentStack)-1]

			if currEl != nil {
				topFromParentStack.TextContent = currEl.TextContent
			}

			if len(parentStack) > 0 {
				parent := parentStack[len(parentStack)-1]
				parent.Children = append(parent.Children, topFromParentStack)
			} else {
				doc.root = topFromParentStack
			}

			currEl = nil
		default: // html.SelfClosingTagToken, html.StartTagToken
			newEl := &Element{TagName: t.Data}

			// ClassName ClassList Id Attributes fields
			if len(t.Attr) > 0 {
				newEl.Attributes = make(attributes)

				for _, a := range t.Attr {
					switch a.Key {
					case "class":
						newEl.ClassName += a.Val
						newEl.ClassList = strings.Split(a.Val, " ")
					case "id":
						newEl.Id = a.Val
					}

					newEl.Attributes[a.Key] = a.Val
				}
			}

			// PreviousElementSibling NextElementSibling fields
			if len(parentStack) > 0 && currEl == nil {
				parent := parentStack[len(parentStack)-1]

				if len(parent.Children) > 0 {
					lastChild := parent.Children[len(parent.Children)-1]
					newEl.PreviousElementSibling = lastChild
					lastChild.NextElementSibling = newEl
				}
			}

			// set ParentElement
			switch {
			case currEl != nil:
				newEl.ParentElement = currEl
			case len(parentStack) > 0:
				newEl.ParentElement = parentStack[len(parentStack)-1]
			}

			switch {
			case t.Type == html.SelfClosingTagToken, t.Type == html.StartTagToken && isSelfClosingTag(t.Data):
				if currEl != nil {
					currEl.Children = append(currEl.Children, newEl)
				} else {
					topFromParentStack := parentStack[len(parentStack)-1]
					topFromParentStack.Children = append(topFromParentStack.Children, newEl)
				}
			case t.Type == html.StartTagToken:
				if currEl != nil {
					parentStack = append(parentStack, currEl)
				}

				currEl = newEl
			}

			switch tName := newEl.TagName; tName {
			case "a":
				doc.Links = append(doc.Links, newEl)
			case "img":
				doc.Images = append(doc.Images, newEl)
			case "body":
				doc.Body = newEl
			case "head":
				doc.Head = newEl
			case "title":
				doc.Title = &newEl.TextContent
			}
		}
	}

	if len(parentStack) != 0 {
		panic("unmatched opening tags. Please, report a bug.")
	}

	return &doc, nil
}
