package dom

import (
	"goDOM/internal/vdom"
)

type Document struct {
	root vdom.Element
}

func CreateDocument(rootEl vdom.Element) Document {
	return Document{rootEl}
}
