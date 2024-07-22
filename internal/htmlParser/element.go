package htmlparser

type element struct {
	tagName    string
	innerHTML  string
	attributes []attribute
	childrens  []element
}

// Parse HTML element from markup. Get custom V-DOM element.
func parseElement(markup string) element {
	var newEl element

	for markupPos, markupChar := range markup {
		if markupChar == rune('>') {
			tag := parseTag(markup[:markupPos+1])

			newEl.tagName = tag.name
			newEl.attributes = tag.attributes

			content := markup[markupPos+1 : len(markup)-len(tag.name)-3]
			contentEndPos := len(content)

			for contentPos, contentChar := range content {
				if contentChar == rune('<') {
					contentEndPos = contentPos
					n := parseElement(content[contentPos:])
					newEl.childrens = append(newEl.childrens, n)

					break
				}
			}

			newEl.innerHTML = content[:contentEndPos]

			break
		}
	}

	return newEl
}
