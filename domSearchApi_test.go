package goDom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type documentTestPair[P any, E *Element | []*Element | bool] struct {
	description string
	params      P
	expect      E
	expectErr   bool
}

var (
	querySelectorTestPair = []documentTestPair[string, *Element]{
		{
			description: "Id query.",
			params:      "#hh",
			expect:      mockEl_h4_2,
		},
		{
			description: "Attribute with value query.",
			params:      "li[data-pull='weee']",
			expect:      mockEl_li_1,
		},
		{
			description: "Multistage query.",
			params:      "#nav_list ul#sub_item_list span",
			expect:      mockEl_span_0,
		},
		{
			description: "Multistage query.",
			params:      "#nav_list ul#sub_item_list",
			expect:      mockEl_ul_2,
		},
		{
			description: "Multistage query.",
			params:      ".bee ul",
			expect:      mockEl_ul_4,
		},
		{
			description: "Not existed element.",
			params:      ".lal",
			expect:      nil,
			expectErr:   true,
		},
	}
	querySelectorAllTestPair = []documentTestPair[string, []*Element]{
		{
			description: "Class query.",
			params:      ".yellow",
			expect:      []*Element{mockEl_li_11, mockEl_li_12},
		},
		{
			description: "Attribute without value query.",
			params:      "button[disabled]",
			expect:      []*Element{mockEl_button_1, mockEl_button_3},
		},
		{
			description: "Multistage query.",
			params:      "footer .button button",
			expect:      []*Element{mockEl_button_3, mockEl_button_4},
		},
		{
			description: "Multi selectors.",
			params:      ".homi, button[disabled], h2",
			expect: []*Element{
				mockEl_address_0,
				mockEl_address_1,
				mockEl_button_1,
				mockEl_button_3,
				mockEl_h2_1,
				mockEl_h2_2,
			},
		},
		// TODO
		// {
		// 	description: "Multi selector with query_operator_all operator.",
		// 	params:      "#sub_item_list *",
		// 	expect: []*Element{
		// 		mockEl_li_8,
		// 		mockEl_li_9,
		// 		mockEl_span_0,
		// 		mockEl_strong_0,
		// 	},
		// },
		{
			description: "Not existed elements.",
			params:      ".lal",
			expect:      nil,
			expectErr:   true,
		},
	}
	getElementByIdTestPair = []documentTestPair[string, *Element]{
		{
			params: "hh",
			expect: mockEl_h4_2,
		},
		{
			description: "Invalid query. Should get error.",
			params:      "hh ll",
			expect:      nil,
			expectErr:   true,
		},
		{
			description: "Not existed element.",
			params:      "lal",
			expect:      nil,
			expectErr:   true,
		},
	}
	getElementsByClassNameTestPair = []documentTestPair[string, []*Element]{
		{
			params: "homi",
			expect: []*Element{mockEl_address_0, mockEl_address_1},
		},
		{
			params:    "hommo",
			expect:    nil,
			expectErr: true,
		},
	}
	getElementsByTagNameTestPair = []documentTestPair[string, []*Element]{
		{
			params:    "lii",
			expect:    nil,
			expectErr: true,
		},
		{
			params: "img",
			expect: []*Element{mockEl_img_1},
		},
		{
			params: "address",
			expect: []*Element{mockEl_address_0, mockEl_address_1},
		},
	}
)

var containsTestPair = []documentTestPair[[]*Element, bool]{
	{
		description: "Contains.",
		params:      []*Element{mockEl_html, mockEl_li_12},
		expect:      true,
	},
	{
		description: "Contains.",
		params:      []*Element{mockEl_head, mockEl_link_1},
		expect:      true,
	},
	{
		description: "Contains.",
		params:      []*Element{mockEl_nav_1, mockEl_li_2},
		expect:      true,
	},
	{
		description: "Doesn't contains.",
		params:      []*Element{mockEl_h4_1, mockEl_h2_1},
		expect:      false,
	},
}

func Test_querySelector(t *testing.T) {
	getValues := func(params string) (*Element, error) {
		return domSearchAPI{}.querySelector(params, mockDOM.root)
	}

	testWithOneResult(t, querySelectorTestPair, getValues)
}

func Test_querySelectorAll(t *testing.T) {
	getValues := func(params string) ([]*Element, error) {
		return domSearchAPI{}.querySelectorAll(params, mockDOM.root)
	}

	testWithFewResults(t, querySelectorAllTestPair, getValues)
}

func Test_getElementById(t *testing.T) {
	getValues := func(params string) (*Element, error) {
		return domSearchAPI{}.getElementById(params, mockDOM.root)
	}

	testWithOneResult(t, getElementByIdTestPair, getValues)
}

func Test_getElementsByClassName(t *testing.T) {
	getValues := func(params string) ([]*Element, error) {
		return domSearchAPI{}.getElementsByClassName(params, mockDOM.root)
	}

	testWithFewResults(t, getElementsByClassNameTestPair, getValues)
}

func Test_getElementsByTagName(t *testing.T) {
	getValues := func(params string) ([]*Element, error) {
		return domSearchAPI{}.getElementsByTagName(params, mockDOM.root)
	}

	testWithFewResults(t, getElementsByTagNameTestPair, getValues)
}

func Test_contains(t *testing.T) {
	for _, p := range containsTestPair {
		v := domSearchAPI{}.contains(p.params[0], p.params[1])

		assert.EqualValues(t, p.expect, v, p.description)
	}
}

func testWithFewResults(t *testing.T, testPairs []documentTestPair[string, []*Element], getValues func(string) ([]*Element, error)) {
	for _, p := range testPairs {
		v, err := getValues(p.params)

		if p.expectErr {
			if err != nil {
				continue
			}

			t.Error("\nfor", p.params, "\nexpected error")
		} else {
			assert.EqualValues(t, p.expect, v, p.description)
		}
	}
}

func testWithOneResult(t *testing.T, testPairs []documentTestPair[string, *Element], getValues func(string) (*Element, error)) {
	for _, p := range testPairs {
		v, err := getValues(p.params)

		if p.expectErr {
			if err != nil {
				continue
			}

			t.Error("\nfor", p.params, "\nexpected error")
		} else {
			assert.EqualValues(t, p.expect, v, p.description)
		}
	}
}
