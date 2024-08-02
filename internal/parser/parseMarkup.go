package parser

import (
	"strings"

	"github.com/bringmetheaugust/goDOM/internal/dom"
)

// Parse markup. Get DOM-like element tree.
func parseMarkup(markup string) *dom.Element {
	var parentStack []dom.Element
	var root dom.Element
	var currEl *dom.Element

	for _, token := range tokenize(markup) {
		switch {
		case strings.HasPrefix(token, "<") && strings.HasSuffix(token, "/>"): // ? self-closing tag (for XHTML)
			tag := parseTag(token[1 : len(token)-2])
			newEl := dom.Element{TagName: tag.name, Attributes: tag.attributes}

			if currEl != nil {
				currEl.Children = append(currEl.Children, newEl)
			} else {
				root = newEl
			}
		case strings.HasPrefix(token, "</"): // tag closes
			if currEl != nil {
				parentStack = append(parentStack, *currEl)
			}

			tagName := token[2 : len(token)-1]

			if len(parentStack) == 0 {
				panic("unmatched closing tag")
			}

			topFromParentStack := &parentStack[len(parentStack)-1]
			parentStack = parentStack[:len(parentStack)-1]

			if topFromParentStack.TagName != tagName {
				panic("mismatched closing tag")
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
		case strings.HasPrefix(token, "<"): // new tag
			tag := parseTag(token[1 : len(token)-1])
			newEl := dom.Element{TagName: tag.name, Attributes: tag.attributes}

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

			if currEl != nil {
				parentStack = append(parentStack, *currEl)
				currEl.Children = append(currEl.Children, newEl)
				newEl.ParentElement = currEl
			}

			currEl = &newEl
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
