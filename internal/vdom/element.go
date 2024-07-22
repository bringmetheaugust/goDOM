package vdom

type Element struct {
	TagName    string
	InnerHTML  string
	Attributes []Attribute
	Childrens  []Element
}
