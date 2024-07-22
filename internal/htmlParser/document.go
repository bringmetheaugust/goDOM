package htmlparser

type Document element

func parseDocument(markup string) Document {
	element := parseElement(markup)

	return Document(element)
}
