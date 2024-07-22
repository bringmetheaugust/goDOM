package htmlparser

// Parsing HTML markup.
// Return browser DOM-like struct.
func Parse(markup string) Document {
	doc := parseDocument(markup)

	return doc
}
