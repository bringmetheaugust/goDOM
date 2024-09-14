package goDom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type domAPIUtilsTestPair[V any, E any] struct {
	target    *Element
	value     V
	expect    E
	expectErr bool
	descr     string
}

var (
	elementMatchesQueryTestPairs = []domAPIUtilsTestPair[string, bool]{
		{
			target: mockEl_header_0,
			value:  "header",
			expect: true,
			descr:  "Should matches by tag name.",
		},
		{
			target: mockEl_ul_1,
			value:  "#nav_list",
			expect: true,
			descr:  "Should matches by Id attribute.",
		},
		{
			target: mockEl_li_1,
			value:  ".red",
			expect: true,
			descr:  "Should matches by class",
		},
		{
			target: mockEl_button_1,
			value:  "button[disabled]",
			expect: true,
			descr:  "Should matches by attribute without value",
		},
		{
			target: mockEl_div_9,
			value:  "div.nested#[data-pic='pow']",
			expect: true,
			descr:  "Should matches by tag name, class, id and attribute",
		},
	}
	findOneByConditionTestPairs = []domAPIUtilsTestPair[func(*Element) bool, any]{
		{
			target: mockEl_div_6,
			value: func(e *Element) bool {
				return e.TagName == "div"
			},
			expect: mockEl_div_7,
			descr:  "Find one nearest element",
		},
		{
			target: mockEl_ul_1,
			value: func(e *Element) bool {
				return e.TagName == "span"
			},
			expect: mockEl_span_0,
			descr:  "Find one far away element",
		},
		{
			target: mockEl_footer_0,
			value: func(e *Element) bool {
				return e.TagName == "strong"
			},
			expectErr: true,
			descr:     "Not found in nested elements.",
		},
	}
	findAllByConditionTestPairs = []domAPIUtilsTestPair[func(*Element) bool, any]{
		{
			target: mockEl_div_5,
			value: func(e *Element) bool {
				return e.TagName == "div"
			},
			expect: []*Element{mockEl_div_6, mockEl_div_7, mockEl_div_8, mockEl_div_9, mockEl_div_10},
			descr:  "Find all nearest element",
		},
		{
			target: mockEl_body_0,
			value: func(e *Element) bool {
				return e.TagName == "button"
			},
			expect: []*Element{mockEl_button_1, mockEl_button_2, mockEl_button_3, mockEl_button_4},
			descr:  "Find all far away element",
		},
		{
			target: mockEl_footer_0,
			value: func(e *Element) bool {
				return e.TagName == "strong"
			},
			expectErr: true,
			descr:     "Not found in nested elements.",
		},
	}
)

func Test_elementMatchesQuery(t *testing.T) {
	for _, p := range elementMatchesQueryTestPairs {
		q, _ := parseQuery(p.value)

		assert.Equal(t, p.expect, domAPIUtils{}.elementMatchesQuery(*q, p.target))
	}
}

func Test_findOneByCondition(t *testing.T) {
	for _, p := range findOneByConditionTestPairs {
		r, err := domAPIUtils{}.findOneByCondition(p.value, p.target)

		if p.expectErr {
			if err != nil {
				continue
			}

			t.Error("\nfor", p.descr, "\nexpected error")
		} else {
			assert.Equal(t, p.expect, r)
		}
	}
}

func Test_findAllByCondition(t *testing.T) {
	for _, p := range findAllByConditionTestPairs {
		r, err := domAPIUtils{}.findAllByCondition(p.value, p.target)

		if p.expectErr {
			if err != nil {
				continue
			}

			t.Error("\nfor", p.descr, "\nexpected error")
		} else {
			assert.Equal(t, p.expect, r)
		}
	}
}
