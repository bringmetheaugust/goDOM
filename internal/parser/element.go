package parser

import (
	"goDOM/internal/dom"
	"strings"
)

// Parse HTML element from markup. Get DOM element.
func parseElement(markup string, parentEl dom.Element) dom.Element {
	var newEl dom.Element

	// if parent is exist
	if parentEl.TagName != "" {
		newEl.ParentElement = &parentEl
	}

	for markupPos, markupChar := range markup {
		if markupChar == rune('>') {
			tag := parseTag(markup[:markupPos+1])

			newEl.TagName = tag.name
			newEl.Attributes = tag.attributes

			// set ClassName, Classlist, Id fields
			for _, attr := range tag.attributes {
				switch true {
				case attr.Name == "class":
					newEl.ClassName = attr.Value
					newEl.ClassList = strings.Fields(attr.Value)
					continue
				case attr.Name == "id":
					newEl.Id = attr.Value
					continue
				}
			}

			content := markup[markupPos+1 : len(markup)-len(tag.name)-3]
			contentEndPos := len(content)

			for contentPos, contentChar := range content {
				if contentChar == rune('<') {
					contentEndPos = contentPos
					childrens := parseElement(content[contentPos:], newEl)
					newEl.Children = append(newEl.Children, childrens)

					if len(newEl.Children) > 0 {
						newEl.FirstChild = &newEl.Children[0]
						newEl.LastChild = &newEl.Children[len(newEl.Children)-1]
					}

					break
				}
			}

			newEl.TextContent = content[:contentEndPos]

			break
		}
	}

	return newEl
}
