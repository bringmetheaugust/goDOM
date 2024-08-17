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
	var parentStack []Element
	var currEl *Element

rootLopp:
	for t := range upStream {
		switch {
		case t.Type == html.CommentToken:
			continue rootLopp
		case t.Type == html.DoctypeToken:
			switch tLow := strings.ToLower(t.Data); {
			case strings.HasPrefix(tLow, "<!doctype html public '-//w3c//dtd xhtml"):
				doc.Doctype = xhtml
			default:
				doc.Doctype = html5
			}
		case t.Type == html.TextToken:
			str := strings.TrimSpace(t.Data)

			if str == "" {
				continue rootLopp
			}

			if currEl != nil {
				currEl.TextContent += str
			} else {
				parentStack[len(parentStack)-1].TextContent += str
			}
		case t.Type == html.EndTagToken:
			if currEl != nil {
				parentStack = append(parentStack, *currEl)
			}

			if len(parentStack) == 0 {
				panic("Error during parsing markup: unmatched closing tag. Please, report a bug.")
			}

			topFromParentStack := &parentStack[len(parentStack)-1]
			parentStack = parentStack[:len(parentStack)-1]

			if topFromParentStack.TagName != t.Data {
				panic("Error during parsing markup: mismatched closing tag. Please, report a bug.")
			}

			if currEl != nil {
				topFromParentStack.TextContent = currEl.TextContent
			}

			if len(parentStack) > 0 {
				parent := &parentStack[len(parentStack)-1]
				parent.Children = append(parent.Children, *topFromParentStack)
			} else {
				doc.root = *topFromParentStack
			}

			currEl = nil
		default: // html.SelfClosingTagToken, html.StartTagToken
			newEl := Element{TagName: t.Data}

			if len(t.Attr) > 0 {
				newEl.Attributes = make(attributes)

				for _, a := range t.Attr {
					v := a.Val

					switch a.Key {
					case "class":
						newEl.ClassName += v
						newEl.ClassList = strings.Split(v, " ")
					case "id":
						newEl.Id = v
					}

					newEl.Attributes[a.Key] = v
				}
			}

			switch {
			case t.Type == html.SelfClosingTagToken, t.Type == html.StartTagToken && isSelfClosingTag(t.Data):
				if currEl != nil {
					currEl.Children = append(currEl.Children, newEl)
				} else {
					topFromParentStack := &parentStack[len(parentStack)-1]
					topFromParentStack.Children = append(topFromParentStack.Children, newEl)
				}
			case t.Type == html.StartTagToken:
				if currEl != nil {
					parentStack = append(parentStack, *currEl)
					currEl.Children = append(currEl.Children, newEl)
					newEl.ParentElement = currEl
				}

				currEl = &newEl
			}

			// *** Document fields
			doc.All = append(doc.All, &newEl)

			// TODO check if add count for optimization

			switch tName := newEl.TagName; tName {
			case "body":
				doc.Body = &newEl
			case "head":
				doc.Head = &newEl
			case "title":
				doc.Title = &newEl.TextContent
			case "a":
				doc.Links = append(doc.Links, &newEl)
			case "img":
				doc.Images = append(doc.Images, &newEl)
			}
		}
	}

	if len(parentStack) != 0 {
		panic("unmatched opening tags. Please, report a bug.")
	}

	return &doc, nil
}
