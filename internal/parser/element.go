package parser

import (
	"goDOM/internal/vdom"
)

// Parse HTML element from markup. Get custom V-DOM element.
func parseElement(markup string) vdom.Element {
	var newEl vdom.Element

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
					newEl.Childrens = append(newEl.Childrens, n)

					break
				}
			}

			newEl.InnerHTML = content[:contentEndPos]

			break
		}
	}

	return newEl
}
