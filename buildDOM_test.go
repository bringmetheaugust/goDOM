package goDom

var (
	mockEl_html = &Element{
		TagName:                "html",
		TextContent:            "",
		Attributes:             attributes{"lang": "en"},
		ClassName:              "",
		ClassList:              nil,
		Id:                     "",
		Children:               children{mockEl_head, mockEl_body_0},
		ParentElement:          nil,
		PreviousElementSibling: nil,
		NextElementSibling:     nil,
	}
	mockEl_head = &Element{
		TagName:     "head",
		TextContent: "",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    children{mockEl_title_1, mockEl_link_1},
	}
	mockEl_title_1 = &Element{
		TagName:     "title",
		TextContent: "Let's try query selector",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    nil,
	}
	mockEl_link_1 = &Element{
		TagName:     "link",
		TextContent: "",
		Attributes:  attributes{"type": "style"},
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    nil,
	}
	mockEl_body_0 = &Element{
		TagName:     "body",
		TextContent: "",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    children{mockEl_header_0, mockEl_main_1, mockEl_footer_1},
	}
	mockEl_header_0 = &Element{
		TagName:     "header",
		TextContent: "",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    children{mockEl_h2_1, mockEl_nav_1, mockEl_div_1},
	}
	mockEl_h2_1 = &Element{
		TagName:     "h2",
		TextContent: "this is header",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    nil,
	}
	mockEl_nav_1 = &Element{
		TagName:     "nav",
		TextContent: "",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    children{mockEl_h3_1, mockEl_ul_1},
	}
	mockEl_h3_1 = &Element{
		TagName:     "h3",
		TextContent: "navigation",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    nil,
	}
	mockEl_ul_1 = &Element{
		TagName:     "ul",
		TextContent: "",
		Attributes:  attributes{"id": "nav_list"},
		ClassName:   "",
		ClassList:   nil,
		Id:          "nav_list",
		Children:    children{mockEl_li_1, mockEl_li_2, mockEl_li_3},
	}
	mockEl_li_1 = &Element{
		TagName:     "li",
		TextContent: "nav item 1",
		Attributes:  attributes{"class": "red", "data-pull": "weee"},
		ClassName:   "red",
		ClassList:   classList{"red"},
		Id:          "",
		Children:    nil,
	}
	mockEl_li_2 = &Element{
		TagName:     "li",
		TextContent: "nav item 2",
		Attributes:  attributes{"class": "red"},
		ClassName:   "red",
		ClassList:   classList{"red"},
		Id:          "",
		Children:    nil,
	}
	mockEl_li_3 = &Element{
		TagName:     "li",
		TextContent: "",
		Attributes:  attributes{"class": "red"},
		ClassName:   "red",
		ClassList:   classList{"red"},
		Id:          "",
		Children:    children{mockEl_h4_1, mockEl_ul_2},
	}
	mockEl_h4_1 = &Element{
		TagName:     "h4",
		TextContent: "nav item 3",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    nil,
	}
	mockEl_ul_2 = &Element{
		TagName:     "ul",
		TextContent: "",
		Attributes:  attributes{"id": "sub_item_list"},
		ClassName:   "",
		ClassList:   nil,
		Id:          "sub_item_list",
		Children:    children{mockEl_li_4, mockEl_li_5},
	}
	mockEl_li_4 = &Element{
		TagName:     "li",
		TextContent: "sub item 1",
		Attributes:  attributes{"class": "white"},
		ClassName:   "white",
		ClassList:   classList{"white"},
		Id:          "",
		Children:    nil,
	}
	mockEl_li_5 = &Element{
		TagName:     "li",
		TextContent: "sub item 2",
		Attributes:  attributes{"class": "white"},
		ClassName:   "white",
		ClassList:   classList{"white"},
		Id:          "",
		Children:    children{mockEl_span_0},
	}
	mockEl_span_0 = &Element{
		TagName:     "span",
		TextContent: "top",
		Attributes:  attributes{"class": "top"},
		ClassName:   "top",
		ClassList:   classList{"top"},
		Id:          "",
		Children:    children{mockEl_strong_0},
	}
	mockEl_strong_0 = &Element{
		TagName:     "strong",
		TextContent: "U are strong!",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    nil,
	}
	mockEl_div_1 = &Element{
		TagName:     "div",
		TextContent: "",
		Attributes:  attributes{"class": "button"},
		ClassName:   "button",
		ClassList:   classList{"button"},
		Id:          "",
		Children:    children{mockEl_button_1, mockEl_button_2},
	}
	mockEl_button_1 = &Element{
		TagName:     "button",
		TextContent: "save",
		Attributes:  attributes{"disabled": "", "data-custom": "shit"},
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    nil,
	}
	mockEl_button_2 = &Element{
		TagName:     "button",
		TextContent: "close",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    nil,
	}
	mockEl_main_1 = &Element{
		TagName:     "main",
		TextContent: "",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    children{mockEl_div_2},
	}
	mockEl_div_2 = &Element{
		TagName:     "div",
		TextContent: "",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    children{mockEl_address_0, mockEl_img_1, mockEl_address_1},
	}
	mockEl_address_0 = &Element{
		TagName:     "address",
		TextContent: "home 1",
		Attributes:  attributes{"class": "homi"},
		ClassName:   "homi",
		ClassList:   classList{"homi"},
		Id:          "",
		Children:    nil,
	}
	mockEl_img_1 = &Element{
		TagName:     "img",
		TextContent: "",
		Attributes:  attributes{"src": "http://zalupa.img.com", "width": "50", "height": "100"},
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    nil,
	}
	mockEl_address_1 = &Element{
		TagName:     "address",
		TextContent: "home 2",
		Attributes:  attributes{"class": "homi homo"},
		ClassName:   "homi homo",
		ClassList:   classList{"homi", "homo"},
		Id:          "",
		Children:    nil,
	}
	mockEl_footer_1 = &Element{
		TagName:     "footer",
		TextContent: "",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    children{mockEl_div_3},
	}
	mockEl_div_3 = &Element{
		TagName:     "div",
		TextContent: "",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    children{mockEl_h2_2, mockEl_nav_2},
	}
	mockEl_h2_2 = &Element{
		TagName:     "h2",
		TextContent: "this is footer",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    nil,
	}
	mockEl_nav_2 = &Element{
		TagName:     "nav",
		TextContent: "",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    children{mockEl_ul_3, mockEl_div_4},
	}
	mockEl_ul_3 = &Element{
		TagName:     "ul",
		TextContent: "",
		Attributes:  attributes{"class": "bee"},
		ClassName:   "bee",
		ClassList:   classList{"bee"},
		Id:          "",
		Children:    children{mockEl_li_6, mockEl_li_7, mockEl_li_10, mockEl_li_11, mockEl_li_12},
	}
	mockEl_li_6 = &Element{
		TagName:     "li",
		TextContent: "nav item 1",
		Attributes:  attributes{"class": "red"},
		ClassName:   "red",
		ClassList:   classList{"red"},
		Id:          "",
		Children:    nil,
	}
	mockEl_li_7 = &Element{
		TagName:     "li",
		TextContent: "",
		Attributes:  attributes{"class": "red"},
		ClassName:   "red",
		ClassList:   classList{"red"},
		Id:          "",
		Children:    children{mockEl_h4_2, mockEl_ul_4},
	}
	mockEl_h4_2 = &Element{
		TagName:     "h4",
		TextContent: "nav item 2",
		Attributes:  attributes{"id": "hh"},
		ClassName:   "",
		ClassList:   nil,
		Id:          "hh",
		Children:    nil,
	}
	mockEl_ul_4 = &Element{
		TagName:     "ul",
		TextContent: "",
		Attributes:  attributes{"class": "bee-bee"},
		ClassName:   "bee-bee",
		ClassList:   classList{"bee-bee"},
		Id:          "",
		Children:    children{mockEl_li_8, mockEl_li_9},
	}
	mockEl_li_8 = &Element{
		TagName:     "li",
		TextContent: "sub item 1",
		Attributes:  attributes{"class": "white"},
		ClassName:   "white",
		ClassList:   classList{"white"},
		Id:          "",
		Children:    nil,
	}
	mockEl_li_9 = &Element{
		TagName:     "li",
		TextContent: "sub item 2",
		Attributes:  attributes{"class": "white"},
		ClassName:   "white",
		ClassList:   classList{"white"},
		Id:          "",
		Children:    nil,
	}
	mockEl_li_10 = &Element{
		TagName:     "li",
		TextContent: "nav item 3",
		Attributes:  attributes{"class": "red"},
		ClassName:   "red",
		ClassList:   classList{"red"},
		Id:          "",
		Children:    nil,
	}
	mockEl_li_11 = &Element{
		TagName:     "li",
		TextContent: "nav item 4",
		Attributes:  attributes{"class": "yellow"},
		ClassName:   "yellow",
		ClassList:   classList{"yellow"},
		Id:          "",
		Children:    nil,
	}
	mockEl_li_12 = &Element{
		TagName:     "li",
		TextContent: "nav item 5",
		Attributes:  attributes{"class": "yellow itt"},
		ClassName:   "yellow itt",
		ClassList:   classList{"yellow", "itt"},
		Id:          "",
		Children:    nil,
	}
	mockEl_div_4 = &Element{
		TagName:     "div",
		TextContent: "",
		Attributes:  attributes{"class": "button"},
		ClassName:   "button",
		ClassList:   classList{"button"},
		Id:          "",
		Children:    children{mockEl_button_3, mockEl_button_4},
	}
	mockEl_button_3 = &Element{
		TagName:     "button",
		TextContent: "delete",
		Attributes:  attributes{"disabled": ""},
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    nil,
	}
	mockEl_button_4 = &Element{
		TagName:     "button",
		TextContent: "close",
		Attributes:  nil,
		ClassName:   "",
		ClassList:   nil,
		Id:          "",
		Children:    nil,
	}
)

var mockDOM = Document{
	Title:   &mockEl_title_1.TextContent,
	Body:    mockEl_body_0,
	Head:    mockEl_header_0,
	Doctype: html5,
	Links:   []*Element{},
	Images:  []*Element{mockEl_img_1},
	root:    buildMockElement(mockEl_html),
}

func buildMockElement(r *Element) *Element {
	for i, c := range r.Children {
		c.ParentElement = r

		if i != 0 {
			c.PreviousElementSibling = r.Children[i-1]
		}

		if i != len(r.Children)-1 {
			c.NextElementSibling = r.Children[i+1]
		}

		buildMockElement(c)
	}

	return r
}

// var (
// 	ignoredTestFields = []string{"domSearchAPI"}
// 	testFilePaths     = []string{"./test/test.html"}
// )

// // Map DOM to ignored some fields for Document and each Element in DOM tree.
// // Cann't add those field to [htmlExpect] variable cause ParentElement is a pointer to parent Element.
// func mapElementForTesting(e *Element) *Element {
// 	mapedStruct, _ := tools.СopyStructWithoutFields[Element](*e, ignoredTestFields)
// 	var childAcc []*Element

// 	for _, child := range mapedStruct.Children {
// 		mapedChild := mapElementForTesting(child)
// 		childAcc = append(childAcc, mapedChild)
// 	}

// 	mapedStruct.Children = childAcc

// 	return &mapedStruct
// }

// func Test_buildDOM(t *testing.T) {
// 	fmt.Printf("\033[33;1m"+"Attention! This test is ignoring %v fields."+"\033[0m"+"\n", ignoredTestFields)

// 	for _, testFilePath := range testFilePaths {
// 		testFile, _ := os.ReadFile(testFilePath)
// 		DOM, _, _ := Create(testFile)

// 		assert.EqualValuesf(t, mockDOM.root, DOM.root, "")
// 	}
// }
