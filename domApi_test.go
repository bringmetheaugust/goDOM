package goDom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type documentTestPair[P any, E *Element | []*Element | bool] struct {
	description string
	target      *Element
	params      P
	expect      E
	expectErr   bool
}

var (
	querySelectorTestPair = []documentTestPair[string, *Element]{
		{
			description: "Tag query.",
			target:      mockEl_div_5,
			params:      "div",
			expect:      mockEl_div_6,
		},
		{
			description: "Id query.",
			target:      mockDOM.root,
			params:      "#hh",
			expect:      mockEl_h4_2,
		},
		{
			description: "Attribute with value query.",
			target:      mockDOM.root,
			params:      "li[data-pull='weee']",
			expect:      mockEl_li_1,
		},
		{
			description: "Multi query.",
			target:      mockDOM.root,
			params:      "#nav_list ul#sub_item_list span",
			expect:      mockEl_span_0,
		},
		{
			description: "Multi query.",
			target:      mockDOM.root,
			params:      "#nav_list ul#sub_item_list",
			expect:      mockEl_ul_2,
		},
		{
			description: "Multi query.",
			target:      mockDOM.root,
			params:      ".bee ul",
			expect:      mockEl_ul_4,
		},
		{
			description: "Not existed element.",
			target:      mockDOM.root,
			params:      ".lal",
			expect:      nil,
			expectErr:   true,
		},
	}
	querySelectorAllTestPair = []documentTestPair[string, []*Element]{
		{
			description: "Class query.",
			target:      mockDOM.root,
			params:      ".yellow",
			expect:      []*Element{mockEl_li_11, mockEl_li_12},
		},
		{
			description: "Attribute without value query.",
			target:      mockDOM.root,
			params:      "button[disabled]",
			expect:      []*Element{mockEl_button_1, mockEl_button_3},
		},
		{
			description: "Multi selectors.",
			target:      mockDOM.root,
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
		{
			description: "Multi selector with query_operator_all operator.",
			target:      mockDOM.root,
			params:      "#sub_item_list *",
			expect: []*Element{
				mockEl_li_4,
				mockEl_li_5,
				mockEl_span_0,
				mockEl_strong_0,
			},
		},
		{
			description: "Not existed elements.",
			target:      mockDOM.root,
			params:      ".lal",
			expect:      nil,
			expectErr:   true,
		},
	}
	getElementByIdTestPair = []documentTestPair[string, *Element]{
		{
			target: mockDOM.root,
			params: "hh",
			expect: mockEl_h4_2,
		},
		{
			target:      mockDOM.root,
			description: "Invalid query. Should get error.",
			params:      "hh ll",
			expect:      nil,
			expectErr:   true,
		},
		{
			target:      mockDOM.root,
			description: "Not existed element.",
			params:      "lal",
			expect:      nil,
			expectErr:   true,
		},
	}
	getElementsByClassNameTestPair = []documentTestPair[string, []*Element]{
		{
			target: mockDOM.root,
			params: "homi",
			expect: []*Element{mockEl_address_0, mockEl_address_1},
		},
		{
			target:    mockDOM.root,
			params:    "hommo",
			expect:    nil,
			expectErr: true,
		},
	}
	getElementsByTagNameTestPair = []documentTestPair[string, []*Element]{
		{
			target:    mockDOM.root,
			params:    "lii",
			expect:    nil,
			expectErr: true,
		},
		{
			target: mockDOM.root,
			params: "img",
			expect: []*Element{mockEl_img_1},
		},
		{
			target: mockDOM.root,
			params: "address",
			expect: []*Element{mockEl_address_0, mockEl_address_1},
		},
	}
)

var containsTestPair = []documentTestPair[[]*Element, bool]{
	{
		description: "Contains.",
		target:      mockDOM.root,
		params:      []*Element{mockEl_html, mockEl_li_12},
		expect:      true,
	},
	{
		description: "Contains.",
		target:      mockDOM.root,
		params:      []*Element{mockEl_head, mockEl_link_1},
		expect:      true,
	},
	{
		description: "Contains.",
		target:      mockDOM.root,
		params:      []*Element{mockEl_nav_1, mockEl_li_2},
		expect:      true,
	},
	{
		description: "Doesn't contains.",
		target:      mockDOM.root,
		params:      []*Element{mockEl_h4_1, mockEl_h2_1},
		expect:      false,
	},
}

func Test_querySelector(t *testing.T) {
	getValues := func(params string, target *Element) (*Element, error) {
		return domAPI{}.querySelector(params, target)
	}

	testWithOneResult(t, querySelectorTestPair, getValues)
}

func Test_querySelectorAll(t *testing.T) {
	getValues := func(params string, target *Element) ([]*Element, error) {
		return domAPI{}.querySelectorAll(params, target)
	}

	testWithFewResults(t, querySelectorAllTestPair, getValues)
}

func Test_getElementById(t *testing.T) {
	getValues := func(params string, target *Element) (*Element, error) {
		return domAPI{}.getElementById(params, target)
	}

	testWithOneResult(t, getElementByIdTestPair, getValues)
}

func Test_getElementsByClassName(t *testing.T) {
	getValues := func(params string, target *Element) ([]*Element, error) {
		return domAPI{}.getElementsByClassName(params, target)
	}

	testWithFewResults(t, getElementsByClassNameTestPair, getValues)
}

func Test_getElementsByTagName(t *testing.T) {
	getValues := func(params string, target *Element) ([]*Element, error) {
		return domAPI{}.getElementsByTagName(params, target)
	}

	testWithFewResults(t, getElementsByTagNameTestPair, getValues)
}

func Test_contains(t *testing.T) {
	for _, p := range containsTestPair {
		v := domAPI{}.contains(p.params[0], p.params[1])

		assert.EqualValues(t, p.expect, v, p.description)
	}
}

func testWithFewResults(
	t *testing.T,
	testPairs []documentTestPair[string, []*Element],
	getValues func(string, *Element) ([]*Element, error),
) {
	for _, p := range testPairs {
		v, err := getValues(p.params, p.target)

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

func testWithOneResult(
	t *testing.T,
	testPairs []documentTestPair[string, *Element],
	getValues func(string, *Element) (*Element, error),
) {
	for _, p := range testPairs {
		v, err := getValues(p.params, p.target)

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
