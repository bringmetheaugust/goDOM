package dom

type Element struct {
	TagName    string
	InnerHTML  string
	Attributes []Attribute
	Children   []Element
}
