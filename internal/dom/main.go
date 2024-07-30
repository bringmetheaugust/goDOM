package dom

// Create new document.
func CreateDocument(rootEl *Element) *Document {
	newDoc := Document{root: *rootEl}

	return &newDoc
}
