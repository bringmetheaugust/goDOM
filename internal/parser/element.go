package parser

import (
	"goDOM/internal/dom"
)

// Parse HTML element from markup. Get DOM element.
func parseElement(markup string) dom.Element {
	var newEl dom.Element

	for markupPos, markupChar := range markup {
		if markupChar == rune('>') {
			tag := parseTag(markup[:markupPos+1])

			newEl.TagName = tag.name
			newEl.Attributes = tag.attributes

			content := markup[markupPos+1 : len(markup)-len(tag.name)-3]
			contentEndPos := len(content)

			for contentPos, contentChar := range content {
				if contentChar == rune('<') {
					contentEndPos = contentPos
					n := parseElement(content[contentPos:])
					newEl.Children = append(newEl.Children, n)

					break
				}
			}

			newEl.InnerHTML = content[:contentEndPos]

			break
		}
	}

	return newEl
}
