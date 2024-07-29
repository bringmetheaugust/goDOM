package parser

// import (
// 	"goDOM/internal/dom"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var htmlTest = ``
// var htmlExpect = dom.Element{
// 	TagName:     "html",
// 	TextContent: "",
// 	Attributes: dom.Attributes{
// 		"lang": "ua",
// 	},
// Children: []dom.Element{
// 	{
// 		TagName:     "li",
// 		TextContent: "ok?",
// 		Attributes: dom.Attributes{
// 			"class": "lol_1",
// 		},
// 		Children:      nil,
// 		ClassName:     "lol_1",
// 		ClassList:     []string{"lol_1"},
// 		FirstChild:    nil,
// 		LastChild:     nil,
// 		Id:            "",
// 		ParentElement: nil,
// 	},
// 	{
// 		TagName:     "li",
// 		TextContent: "ok???",
// 		Attributes: dom.Attributes{
// 			"class": "lol_2",
// 		},
// 		Children:      nil,
// 		ClassName:     "lol_2",
// 		ClassList:     []string{"lol_2"},
// 		FirstChild:    nil,
// 		LastChild:     nil,
// 		Id:            "",
// 		ParentElement: nil,
// 	},
// 	{
// 		TagName:     "li",
// 		TextContent: "ok",
// 		Attributes: dom.Attributes{
// 			"class": "lol_3",
// 		},
// 		Children:      nil,
// 		ClassName:     "lol_3",
// 		ClassList:     []string{"lol_3"},
// 		FirstChild:    nil,
// 		LastChild:     nil,
// 		Id:            "",
// 		ParentElement: nil,
// 	},
// },
// }

// func Test_ParseHTML(t *testing.T) {
// 	v, _ := ParseHTML(htmlTest)

// 	assert.EqualValuesf(t, htmlExpect, v, "")
// }
