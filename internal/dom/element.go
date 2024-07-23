package dom

type Element struct {
	TagName       string
	TextContent   string // only own text/content
	Attributes    []Attribute
	Children      []Element
	ClassName     string
	ClassList     []string
	FirstChild    *Element
	LastChild     *Element
	Id            string
	ParentElement *Element
}
