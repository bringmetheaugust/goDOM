package goDom

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	"github.com/bringmetheaugust/goDOM/tools"
// 	"github.com/stretchr/testify/assert"
// )

// var htmlExpect = Document{
// 	Links: []*Element{
// 		{
// 			TagName:     "a",
// 			Attributes:  attributes{"href": "nahuy"},
// 			TextContent: "to army",
// 		},
// 	},
// 	Images: []*Element{
// 		{
// 			TagName:    "img",
// 			Attributes: attributes{"src": "https://hell.com"},
// 			Children:   nil,
// 		},
// 		{
// 			TagName:    "img",
// 			Attributes: attributes{"src": "https://heaven.com"},
// 			Children:   nil,
// 		},
// 	},
// 	root: Element{
// 		TagName:    "html",
// 		Attributes: attributes{"lang": "en"},
// 		Children: children{
// 			{
// 				TagName: "head",
// 				Children: children{
// 					{
// 						TagName:     "title",
// 						TextContent: "Test DOM",
// 					},
// 					{
// 						TagName:    "meta",
// 						Attributes: attributes{"charset": "UTF-8"},
// 					},
// 					{
// 						TagName:    "link",
// 						Attributes: attributes{"href": "ururu", "id": "lolipop"},
// 						Id:         "lolipop",
// 					},
// 				},
// 			},
// 			{
// 				TagName: "body",
// 				Children: children{
// 					{
// 						TagName:     "a",
// 						Attributes:  attributes{"href": "nahuy"},
// 						TextContent: "to army",
// 					},
// 					{
// 						TagName:    "ul",
// 						Attributes: attributes{"class": "lol lul", "id": "ou"},
// 						Id:         "ou",
// 						ClassName:  "lol lul",
// 						ClassList:  classList{"lol", "lul"},
// 						Children: children{
// 							{
// 								TagName:     "li",
// 								Attributes:  attributes{"id": "ouu"},
// 								Id:          "ouu",
// 								TextContent: "li 0",
// 								Children: children{
// 									{
// 										TagName:    "ul",
// 										Attributes: attributes{"class": "two"},
// 										ClassName:  "two",
// 										ClassList:  classList{"two"},
// 										Children: children{
// 											{
// 												TagName:     "li",
// 												Attributes:  attributes{"href": "afa sada_1"},
// 												TextContent: "li 1",
// 												Children: children{
// 													{
// 														TagName:     "span",
// 														TextContent: "ahaha from li 1",
// 														Children:    nil,
// 													},
// 												},
// 											},
// 											{
// 												TagName:     "li",
// 												Attributes:  attributes{"href": "afa sada_2"},
// 												TextContent: "li 2",
// 												Children: children{
// 													{
// 														TagName:     "span",
// 														TextContent: "ahaha from li 2",
// 														Children:    nil,
// 													},
// 												},
// 											},
// 											{
// 												TagName:     "li",
// 												Attributes:  attributes{"href": "afa sada_3"},
// 												TextContent: "li 3li 3",
// 												Children: children{
// 													{
// 														TagName:     "span",
// 														TextContent: "ahaha from li 3",
// 														Children:    nil,
// 													},
// 													{
// 														TagName:    "img",
// 														Attributes: attributes{"src": "https://hell.com"},
// 														Children:   nil,
// 													},
// 												},
// 											},
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 					{
// 						TagName:     "button",
// 						TextContent: "mm?",
// 						Attributes:  attributes{"disabled": ""},
// 						Children:    nil,
// 					},
// 					{
// 						TagName:    "img",
// 						Attributes: attributes{"src": "https://heaven.com"},
// 						Children:   nil,
// 					},
// 					{
// 						TagName:     "span",
// 						TextContent: "this is span, baby",
// 						Children:    nil,
// 					},
// 				},
// 			},
// 		},
// 	},
// }

// var ignoredTestFields = []string{"ParentElement", "domSearchAPI"}

// // var testFilePaths = []string{"./test/parse_markup_html5.html", "./test/parse_markup_xhtml.html"}
// var testFilePaths = []string{"./test/parse_markup_html5.html"}

// // Map DOM to ignored some fields for Document and each Element in DOM tree.
// // Cann't add those field to [htmlExpect] variable cause ParentElement is a pointer to parent Element.
// func mapDomForTesting(DOM *Document) *Document {
// 	mapedStruct, _ := tools.СopyStructWithoutFields[Element](*DOM, ignoredTestFields)
// 	var childAcc []Element

// 	for _, child := range mapedStruct.Children {
// 		mapedChild := mapDomForTesting(&child)
// 		childAcc = append(childAcc, *mapedChild)
// 	}

// 	mapedStruct.Children = childAcc

// 	return &mapedStruct
// }

// func Test_parseMarkup(t *testing.T) {
// 	fmt.Printf("\033[33;1m"+"Attention! This test is ignoring %v fields."+"\033[0m"+"\n", ignoredTestFields)

// 	for _, testFilePath := range testFilePaths {
// 		testFile, _ := os.ReadFile(testFilePath)
// 		DOM, _ := Create(testFile)
// 		mapedDOM := mapDomForTesting(DOM) // map Elemenet struct and remove some fields

// 		assert.EqualValuesf(t, &htmlExpect, mapedDOM, "")
// 	}
// }
