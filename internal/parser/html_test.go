package parser

import (
	"goDOM/internal/dom"
	"goDOM/tools"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var htmlExpect = &dom.Element{
	TagName:    "html",
	Attributes: dom.Attributes{"lang": "en"},
	Children: []dom.Element{
		{
			TagName: "body",
			Children: []dom.Element{
				{
					TagName:    "ul",
					Attributes: dom.Attributes{"class": "lol lul", "id": "ou"},
					Id:         "ou",
					ClassName:  "lol lul",
					ClassList:  []string{"lol", "lul"},
					Children: []dom.Element{
						{
							TagName:     "li",
							Attributes:  dom.Attributes{"id": "ouu"},
							Id:          "ouu",
							TextContent: "li 0",
							Children: []dom.Element{
								{
									TagName:    "ul",
									Attributes: dom.Attributes{"class": "two"},
									ClassName:  "two",
									ClassList:  []string{"two"},
									Children: []dom.Element{
										{
											TagName:     "li",
											Attributes:  dom.Attributes{"href": "afa sada_1"},
											TextContent: "li 1",
											Children:    nil,
										},
										{
											TagName:     "li",
											Attributes:  dom.Attributes{"href": "afa sada_2"},
											TextContent: "li 2",
											Children:    nil,
										},
										{
											TagName:     "li",
											Attributes:  dom.Attributes{"href": "afa sada_3"},
											TextContent: "li 3",
											Children:    nil,
										},
									},
								},
							},
						},
					},
				},
				{
					TagName:     "span",
					TextContent: "this is span, baby",
					Children:    nil,
				},
			},
		},
	},
}

func mapDom(DOM *dom.Element) interface{} {
	newStruct := tools.CopyStructWithoutFields(DOM, []string{"ParentElement"}).(dom.Element)

	for _, child := range newStruct.Children {
		child = mapDom(&child).(dom.Element)
	}

	return newStruct
}

// TODO: temporary frozen
func Tessssssssssssst_ParseHTML(t *testing.T) {
	testFile, _ := os.ReadFile("../../test/parse_html.html")
	DOM, _ := ParseHTML(string(testFile))

	// ? Map Elemenet struct and remove ParentElement field.
	mapedDOM := mapDom(DOM)

	assert.EqualValuesf(t, htmlExpect, mapedDOM, "")
}
