package parser

import (
	"fmt"
	"os"
	"testing"

	"github.com/bringmetheaugust/goDOM/internal/dom"
	"github.com/bringmetheaugust/goDOM/tools"
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
											Children: []dom.Element{
												{
													TagName:     "span",
													TextContent: "ahaha from li 1",
													Children:    nil,
												},
											},
										},
										{
											TagName:     "li",
											Attributes:  dom.Attributes{"href": "afa sada_2"},
											TextContent: "li 2",
											Children: []dom.Element{
												{
													TagName:     "span",
													TextContent: "ahaha from li 2",
													Children:    nil,
												},
											},
										},
										{
											TagName:     "li",
											Attributes:  dom.Attributes{"href": "afa sada_3"},
											TextContent: "li 3li 3",
											Children: []dom.Element{
												{
													TagName:     "span",
													TextContent: "ahaha from li 3",
													Children:    nil,
												},
											},
										},
									},
								},
							},
						},
					},
				},
				{
					TagName:     "button",
					TextContent: "mm?",
					Attributes:  dom.Attributes{"disabled": ""},
					Children:    nil,
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
var ignoredTestFields = []string{"ParentElement"}

// Remove ParentElement field from each Element in DOM tree.
// Cann't add this field to [htmlExpect] variable cause ParentElement is a pointer to parent Element.
func mapDomForTesting(DOM *dom.Element) *dom.Element {
	mapedStruct, _ := tools.СopyStructWithoutFields[dom.Element](*DOM, ignoredTestFields)
	var childAcc []dom.Element

	for _, child := range mapedStruct.Children {
		mapedChild := mapDomForTesting(&child)
		childAcc = append(childAcc, *mapedChild)
	}

	mapedStruct.Children = childAcc

	return &mapedStruct
}

func Test_parseMarkup(t *testing.T) {
	fmt.Printf("\033[33;1m"+"Attention! This test is ignoring %v fields."+"\033[0m"+"\n", ignoredTestFields)

	testFile, _ := os.ReadFile("../../test/parse_markup.html")
	DOM, _ := Parse(string(testFile))
	mapedDOM := mapDomForTesting(DOM) // map Elemenet struct and remove some fields

	assert.EqualValuesf(t, htmlExpect, mapedDOM, "")
}
