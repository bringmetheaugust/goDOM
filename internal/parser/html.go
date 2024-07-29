package parser

import (
	"goDOM/internal/dom"
	"goDOM/internal/errors"
	"strings"
)

// TODO: parent Tag, FirstChild, LastChild
// Parse HTML from markup. Get DOM tree.
func ParseHTML(markup string) (dom.Element, error) {
	markup = strings.ReplaceAll(markup, `"`, `'`) // cause string() shielding /"
	markup = strings.TrimSpace(markup)

	if len(markup) == 0 {
		return dom.Element{}, errors.InvalidRequest{Place: "ParseHTML"}
	}

	var parentStack []dom.Element
	var root dom.Element
	var currEl *dom.Element

	for _, token := range tokenize(markup) {
		switch {
		// ? Self-closing tag (for XHTML)
		case strings.HasPrefix(token, "<") && strings.HasSuffix(token, "/>"):
			tag := parseTag(token[1 : len(token)-2])
			newEl := dom.Element{TagName: tag.name, Attributes: tag.attributes}

			if currEl != nil {
				currEl.Children = append(currEl.Children, newEl)
			} else {
				root = newEl
			}
		case strings.HasPrefix(token, "</"): // Closing tag
			if currEl != nil {
				parentStack = append(parentStack, *currEl)
			}

			tagName := token[2 : len(token)-1]

			if len(parentStack) == 0 {
				panic("unmatched closing tag")
			}

			top := parentStack[len(parentStack)-1]
			parentStack = parentStack[:len(parentStack)-1]

			if top.TagName != tagName {
				panic("mismatched closing tag")
			}

			if currEl != nil {
				top.TextContent = currEl.TextContent
			}

			if len(parentStack) > 0 {
				parent := &parentStack[len(parentStack)-1]
				parent.Children = append(parent.Children, top)
			} else {
				root = top
			}

			currEl = nil
		case strings.HasPrefix(token, "<"): // Opening tag
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

	return root, nil
}

func tokenize(input string) []string {
	var tokens []string
	var currentToken strings.Builder

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '<':
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}

			currentToken.WriteByte('<')
		case '>':
			currentToken.WriteByte('>')
			tokens = append(tokens, currentToken.String())
			currentToken.Reset()
		default:
			currentToken.WriteByte(input[i])
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}
