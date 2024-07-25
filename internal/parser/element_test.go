package parser

import (
	"goDOM/internal/dom"
	"testing"

	"github.com/stretchr/testify/assert"
)

type elementArgs struct {
	markup        string
	parentElement dom.Element
}
type elementTestPair struct {
	value     elementArgs
	expect    dom.Element
	expectErr bool
}

var elementTests = []elementTestPair{
	{
		elementArgs{
			"<div class='lol lil' id='cool' hidden></div>",
			dom.Element{},
		},
		dom.Element{
			TagName:     "div",
			TextContent: "",
			Attributes: []dom.Attribute{
				{Name: "class", Value: "lol lil"},
				{Name: "id", Value: "cool"},
				{Name: "hidden", Value: ""},
			},
			Children:      nil,
			ClassName:     "lol lil",
			ClassList:     []string{"lol", "lil"},
			FirstChild:    nil,
			LastChild:     nil,
			Id:            "cool",
			ParentElement: nil,
		},
		false,
	},
	{
		elementArgs{
			"<a href='/usa' style='color: red; display: block;' data-src='Obama'>Biden</a>",
			dom.Element{},
		},
		dom.Element{
			TagName:     "a",
			TextContent: "Biden",
			Attributes: []dom.Attribute{
				{Name: "href", Value: "/usa"},
				{Name: "style", Value: "color: red; display: block;"},
				{Name: "data-src", Value: "Obama"},
			},
			Children:      nil,
			ClassName:     "",
			ClassList:     nil,
			FirstChild:    nil,
			LastChild:     nil,
			Id:            "",
			ParentElement: nil,
		},
		false,
	},
	// TODO
	// {
	// 	elementArgs{
	// 		`<ul id='pizdets'>
	// 			<li class='lol_1'>ok?</li>
	// 			<li class='lol_2'>ok???</li>
	// 			<li class='lol_3'>ok</li>
	// 		</ul>`,
	// 		dom.Element{},
	// 	},
	// 	dom.Element{
	// 		TagName:     "ul",
	// 		TextContent: "",
	// 		Attributes: []dom.Attribute{
	// 			{Name: "id", Value: "pizdets"},
	// 		},
	// 		Children: []dom.Element{
	// 			{
	// 				TagName:     "li",
	// 				TextContent: "ok?",
	// 				Attributes: []dom.Attribute{
	// 					{Name: "class", Value: "lol_1"},
	// 				},
	// 				Children:      nil,
	// 				ClassName:     "lol_1",
	// 				ClassList:     []string{"lol_1"},
	// 				FirstChild:    nil,
	// 				LastChild:     nil,
	// 				Id:            "",
	// 				ParentElement: nil,
	// 			},
	// 			{
	// 				TagName:     "li",
	// 				TextContent: "ok???",
	// 				Attributes: []dom.Attribute{
	// 					{Name: "class", Value: "lol_2"},
	// 				},
	// 				Children:      nil,
	// 				ClassName:     "lol_2",
	// 				ClassList:     []string{"lol_2"},
	// 				FirstChild:    nil,
	// 				LastChild:     nil,
	// 				Id:            "",
	// 				ParentElement: nil,
	// 			},
	// 			{
	// 				TagName:     "li",
	// 				TextContent: "ok",
	// 				Attributes: []dom.Attribute{
	// 					{Name: "class", Value: "lol_3"},
	// 				},
	// 				Children:      nil,
	// 				ClassName:     "lol_3",
	// 				ClassList:     []string{"lol_3"},
	// 				FirstChild:    nil,
	// 				LastChild:     nil,
	// 				Id:            "",
	// 				ParentElement: nil,
	// 			},
	// 		},
	// 		ClassName: "",
	// 		ClassList: nil,
	// 		FirstChild: &dom.Element{
	// 			TagName:     "li",
	// 			TextContent: "ok?",
	// 			Attributes: []dom.Attribute{
	// 				{Name: "class", Value: "lol_1"},
	// 			},
	// 			Children:      nil,
	// 			ClassName:     "lol_1",
	// 			ClassList:     []string{"lol_1"},
	// 			FirstChild:    nil,
	// 			LastChild:     nil,
	// 			Id:            "",
	// 			ParentElement: nil,
	// 		},
	// 		LastChild: &dom.Element{
	// 			TagName:     "li",
	// 			TextContent: "ok",
	// 			Attributes: []dom.Attribute{
	// 				{Name: "class", Value: "lol_3"},
	// 			},
	// 			Children:      nil,
	// 			ClassName:     "lol_3",
	// 			ClassList:     []string{"lol_3"},
	// 			FirstChild:    nil,
	// 			LastChild:     nil,
	// 			Id:            "",
	// 			ParentElement: nil,
	// 		},
	// 		Id:            "",
	// 		ParentElement: nil,
	// 	},
	// 	false,
	// },
}

func Test_parseElement(t *testing.T) {
	for _, pair := range elementTests {
		v := parseElement(pair.value.markup, pair.value.parentElement)

		assert.EqualValuesf(t, pair.expect, v, "")
	}
}
