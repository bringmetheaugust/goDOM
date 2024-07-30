package parser

import (
	"goDOM/internal/dom"
	"goDOM/internal/errors"
	"strings"
)

// Parse HTML from markup. Get DOM tree.
func ParseHTML(markup string) (*dom.Element, error) {
	markup = normalize(markup)

	if len(markup) == 0 {
		return &dom.Element{}, errors.InvalidRequest{Place: "ParseHTML"}
	}

	var parentStack []dom.Element
	var root dom.Element
	var currEl *dom.Element

	for _, token := range tokenize(markup) {
		switch {
		case strings.HasPrefix(token, "<") && strings.HasSuffix(token, "/>"): // ? Self-closing tag (for XHTML)
			tag := parseTag(token[1 : len(token)-2])
			newEl := dom.Element{TagName: tag.name, Attributes: tag.attributes}

			if currEl != nil {
				currEl.Children = append(currEl.Children, newEl)
			} else {
				root = newEl
			}
		case strings.HasPrefix(token, "</"): // Tag closes
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
		case strings.HasPrefix(token, "<"): // New tag
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
			}
		}
	}

	if len(parentStack) != 0 {
		panic("unmatched opening tags")
	}

	return &root, nil
}
