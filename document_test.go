package goDom

import (
	"os"
	"testing"

	"github.com/bringmetheaugust/goDOM/tools"
	"github.com/stretchr/testify/assert"
)

type documentTestPair struct {
	description string
	params      string
	expect      any
	expectErr   bool
}

var testFile, _ = os.ReadFile("./test/search.html")
var document, _ = Create(testFile)
var documnetIgnoredTestFields = []string{"ParentElement", "Children"}

var querySelectorTestPair = []documentTestPair{
	{
		description: "Multistage query.",
		params:      "#nav_list ul#sub_item_list span",
		expect: Element{
			TagName:     "span",
			TextContent: "top",
			Attributes:  attributes{"class": "top"},
			ClassName:   "top",
			ClassList:   []string{"top"},
		},
		expectErr: false,
	},
	{
		description: "Multistage query.",
		params:      "#nav_list ul#sub_item_list",
		expect: Element{
			TagName:     "ul",
			Id:          "sub_item_list",
			TextContent: "",
			Attributes:  attributes{"id": "sub_item_list"},
		},
		expectErr: false,
	},
	{
		description: "Id query.",
		params:      "#hh",
		expect: Element{
			TagName:     "h4",
			Id:          "hh",
			TextContent: "nav item 2",
			Attributes:  attributes{"id": "hh"},
		},
		expectErr: false,
	},
	{
		description: "Not existed element.",
		params:      ".lal",
		expect:      Element{},
		expectErr:   true,
	},
}

var querySelectorAllTestPair = []documentTestPair{
	{
		description: "Multystage query.",
		params:      "footer .button button",
		expect: []Element{
			{
				TagName:     "button",
				TextContent: "delete",
				Attributes:  attributes{"disabled": ""},
			},
			{
				TagName:     "button",
				TextContent: "close",
				Attributes:  nil,
			},
		},
		expectErr: false,
	},
	{
		description: "Class query.",
		params:      ".yellow",
		expect: []Element{
			{
				TagName:     "li",
				TextContent: "nav item 4",
				Attributes:  attributes{"class": "yellow"},
				ClassName:   "yellow",
				ClassList:   []string{"yellow"},
			},
			{
				TagName:     "li",
				TextContent: "nav item 5",
				Attributes:  attributes{"class": "yellow itt"},
				ClassName:   "yellow itt",
				ClassList:   []string{"yellow", "itt"},
			},
		},
		expectErr: false,
	},
	{
		description: "Not existed elements.",
		params:      ".lal",
		expect:      nil,
		expectErr:   true,
	},
}

var getElementByIdTestPair = []documentTestPair{
	{
		params: "hh",
		expect: Element{
			TagName:     "h4",
			Id:          "hh",
			TextContent: "nav item 2",
			Attributes:  attributes{"id": "hh"},
		},
		expectErr: false,
	},
	{
		params:    "hh ll",
		expect:    Element{},
		expectErr: true,
	},
	{
		description: "Not existed element.",
		params:      "lal",
		expect:      Element{},
		expectErr:   true,
	},
}

var getElementsByClassNameTestPair = []documentTestPair{
	{
		params: "homi",
		expect: []Element{
			{
				TagName:     "address",
				TextContent: "home 1",
				Attributes:  attributes{"class": "homi"},
				ClassName:   "homi",
				ClassList:   []string{"homi"},
			},
			{
				TagName:     "address",
				TextContent: "home 2",
				Attributes:  attributes{"class": "homi homo"},
				ClassName:   "homi homo",
				ClassList:   []string{"homi", "homo"},
			},
		},
		expectErr: false,
	},
	{
		params:    "hommo",
		expect:    nil,
		expectErr: true,
	},
}

var getElementsByTagNameTestPair = []documentTestPair{
	{
		params: "address",
		expect: []Element{
			{
				TagName:     "address",
				TextContent: "home 1",
				Attributes:  attributes{"class": "homi"},
				ClassName:   "homi",
				ClassList:   []string{"homi"},
			},
			{
				TagName:     "address",
				TextContent: "home 2",
				Attributes:  attributes{"class": "homi homo"},
				ClassName:   "homi homo",
				ClassList:   []string{"homi", "homo"},
			},
		},
		expectErr: false,
	},
	{
		params: "img",
		expect: []Element{
			{
				TagName:    "img",
				Attributes: attributes{"src": "http://zalupa.img.com", "width": "50", "height": "100"},
			},
		},
		expectErr: false,
	},
	{
		params:    "lii",
		expect:    nil,
		expectErr: true,
	},
}

func Test_QuerySelector(t *testing.T) {
	getValues := func(params string) (Element, error) {
		return document.QuerySelector(params)
	}

	testWithOneResult(t, querySelectorTestPair, getValues)
}

func Test_QuerySelectorAll(t *testing.T) {
	getValues := func(params string) ([]Element, error) {
		return document.QuerySelectorAll(params)
	}

	testWithFewResults(t, querySelectorAllTestPair, getValues)
}

func Test_GetElementById(t *testing.T) {
	getValues := func(params string) (Element, error) {
		return document.GetElementById(params)
	}

	testWithOneResult(t, getElementByIdTestPair, getValues)
}

func Test_GetElementsByClassName(t *testing.T) {
	getValues := func(params string) ([]Element, error) {
		return document.GetElementsByClassName(params)
	}

	testWithFewResults(t, getElementsByClassNameTestPair, getValues)
}

func Test_GetElementsByTagName(t *testing.T) {
	getValues := func(params string) ([]Element, error) {
		return document.GetElementsByTagName(params)
	}

	testWithFewResults(t, getElementsByTagNameTestPair, getValues)
}

func testWithFewResults(t *testing.T, testPairs []documentTestPair, getValues func(string) ([]Element, error)) {
	for _, pair := range testPairs {
		v, err := getValues(pair.params)
		var vMaped []Element

		for _, el := range v {
			mapedV, _ := tools.СopyStructWithoutFields[Element](el, documnetIgnoredTestFields)
			vMaped = append(vMaped, mapedV)
		}

		if pair.expectErr {
			if err != nil {
				continue
			}

			t.Error("\nfor", pair.params, "\nexpected error")
		} else {
			assert.EqualValuesf(t, pair.expect, vMaped, pair.description)
		}
	}
}

func testWithOneResult(t *testing.T, testPairs []documentTestPair, getValues func(string) (Element, error)) {
	for _, pair := range testPairs {
		v, err := getValues(pair.params)
		vMaped, _ := tools.СopyStructWithoutFields[Element](v, documnetIgnoredTestFields)

		if pair.expectErr {
			if err != nil {
				continue
			}

			t.Error("\nfor", pair.params, "\nexpected error")
		} else {
			assert.EqualValuesf(t, pair.expect, vMaped, pair.description)
		}
	}
}
