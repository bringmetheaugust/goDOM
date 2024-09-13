package goDom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type JQTestPair[V any, E any] struct {
	target    JQ
	value     V
	expect    E
	expectErr bool
	desrc     string
}

var (
	attrTestPairs = []JQTestPair[string, string]{
		{
			target: *JQ{}.new(mockEl_li_1),
			value:  "data-pull",
			expect: "weee",
			desrc:  "Existed attribute.",
		},
		{
			target:    *JQ{}.new(mockEl_li_5),
			value:     "no-existed",
			expect:    "",
			expectErr: true,
			desrc:     "Attribute doesn't exist.",
		},
	}
	childrenTestPairs = []JQTestPair[[]string, JQ]{
		{
			target: *JQ{}.new(mockEl_ul_3),
			expect: *JQ{}.new(mockEl_li_6, mockEl_li_7, mockEl_li_10, mockEl_li_11, mockEl_li_12),
			desrc:  "Get all childrens.",
		},
		{
			target: *JQ{}.new(mockEl_strong_0),
			expect: *JQ{}.new(),
			desrc:  "No childrens.",
		},
		{
			target: *JQ{}.new(mockEl_ul_3, mockEl_div_2, mockEl_ul_2, mockEl_div_1),
			value:  []string{".itt, .homo"},
			expect: *JQ{}.new(mockEl_li_12, mockEl_address_1),
			desrc:  "Get all childrens by multi query.",
		},
		{
			target: *JQ{}.new(mockEl_div_1, mockEl_ul_2, mockEl_main_1),
			value:  []string{"footer, #nav_list"},
			expect: *JQ{}.new(),
			desrc:  "No childrens by multi query.",
		},
	}
	filterTestPairs = []JQTestPair[any, JQ]{
		{
			target: *JQ{}.new(mockEl_span_0, mockEl_li_2, mockEl_li_3, mockEl_ul_2, mockEl_button_1),
			value:  "li.red",
			expect: *JQ{}.new(mockEl_li_2, mockEl_li_3),
			desrc:  "Filter jQuery elements by query.",
		},
		{
			target: *JQ{}.new(mockEl_span_0, mockEl_li_2, mockEl_li_3, mockEl_ul_2, mockEl_button_1),
			value:  ".top, li.red",
			expect: *JQ{}.new(mockEl_span_0, mockEl_li_2, mockEl_li_3),
			desrc:  "Filter jQuery elements by multi query.",
		},
		{
			target: *JQ{}.new(mockEl_span_0, mockEl_li_2, mockEl_li_3, mockEl_ul_2, mockEl_button_1),
			value:  ".rd",
			expect: *JQ{}.new(),
			desrc:  "Filter jQuery elements by non-existed query.",
		},
		{
			target: *JQ{}.new(mockEl_footer_0, mockEl_ul_3, mockEl_button_1, mockEl_ul_4, mockEl_address_1),
			value: func(i *JQ) bool {
				attr, _ := i.Attr("class")

				return attr == "bee"
			},
			expect: *JQ{}.new(mockEl_ul_3),
			desrc:  "Filter jQuery elements by callback function.",
		},
	}
	findTestPairs = []JQTestPair[string, JQ]{
		{
			target: *JQ{}.new(mockEl_nav_1),
			value:  ".white",
			expect: *JQ{}.new(mockEl_li_4, mockEl_li_5),
			desrc:  "Find jQuery elements by query.",
		},
		{
			target: *JQ{}.new(mockEl_footer_0, mockEl_div_1, mockEl_main_1, mockEl_ul_1),
			value:  ".yellow, .homo",
			expect: *JQ{}.new(mockEl_li_11, mockEl_li_12, mockEl_address_1),
			desrc:  "Filter jQuery elements by multi queries.",
		},
		{
			target: *JQ{}.new(mockEl_nav_1),
			value:  ".my-ass",
			expect: *JQ{}.new(),
			desrc:  "Find jQuery elements by non-existed query.",
		},
	}
	firstTestPairs = []JQTestPair[*JQ, JQ]{
		{
			target: *JQ{}.new(mockEl_button_1, mockEl_button_2),
			expect: *JQ{}.new(mockEl_button_1),
			desrc:  "First existed jQuery element.",
		},
		{
			target: *JQ{}.new(),
			expect: *JQ{}.new(),
			desrc:  "Non-existed first jQuery element.",
		},
	}
	hasTestPairs = []JQTestPair[any, JQ]{
		{
			target: *JQ{}.new(mockEl_li_1, mockEl_li_2, mockEl_li_3, mockEl_li_6, mockEl_li_7),
			value:  ".bee-bee",
			expect: *JQ{}.new(mockEl_li_7),
			desrc:  "Should has element by query.",
		},
		{
			target: *JQ{}.new(mockEl_li_1, mockEl_li_2, mockEl_li_3, mockEl_li_6, mockEl_li_7),
			value:  "#nav_list",
			expect: *JQ{}.new(),
			desrc:  "Should not has element by query.",
		},
		{
			target: *JQ{}.new(mockEl_li_1, mockEl_li_2, mockEl_li_3, mockEl_li_6, mockEl_li_7),
			value:  JQ{}.new(mockEl_ul_4),
			expect: *JQ{}.new(mockEl_li_7),
			desrc:  "Should has element by JQ element.",
		},
		{
			target: *JQ{}.new(mockEl_li_1, mockEl_li_2, mockEl_li_3, mockEl_li_6, mockEl_li_7),
			value:  JQ{}.new(mockEl_footer_0),
			expect: *JQ{}.new(),
			desrc:  "Should not has element by JQ element.",
		},
	}
	hasCLassTestPairs = []JQTestPair[string, bool]{
		{
			target: *JQ{}.new(mockEl_address_0),
			value:  "homi",
			expect: true,
			desrc:  "Contains class.",
		},
		{
			target: *JQ{}.new(mockEl_address_0),
			value:  "homo",
			expect: false,
			desrc:  "Doesn't contain class.",
		},
		{
			target: *JQ{}.new(),
			value:  ".bee",
			expect: false,
			desrc:  "Doesn't have class in empty query.",
		},
	}
)

func TestAttr(t *testing.T) {
	for _, p := range attrTestPairs {
		r, err := p.target.Attr(p.value)

		if p.expectErr {
			if err != nil {
				continue
			}

			t.Error("\nfor", p.value, "\nexpected error")
		} else {
			assert.EqualValues(t, p.expect, r, p.desrc)
		}
	}
}

func TestChildren(t *testing.T) {
	for _, p := range childrenTestPairs {
		r := p.target.Children(p.value...)

		assert.EqualValues(t, p.expect, r, p.desrc)
	}
}

func TestEach(t *testing.T) {
	var expectedCollector = []string{mockEl_li_4.ClassName, mockEl_address_0.ClassName, mockEl_ul_4.ClassName}
	var collector []string

	JQ{}.new(mockEl_li_4, mockEl_address_0, mockEl_ul_4).Each(func(j *JQ) {
		attr, err := j.Attr("class")

		if err != nil {
			return
		}

		collector = append(collector, attr)
	})

	assert.EqualValues(t, expectedCollector, collector, "Collect array by `each` callback.")
}

func TestFilter(t *testing.T) {
	for _, p := range filterTestPairs {
		r := p.target.Filter(p.value)

		assert.EqualValues(t, p.expect, r, p.desrc)
	}
}

func TestFind(t *testing.T) {
	for _, p := range findTestPairs {
		r := p.target.Find(p.value)

		assert.EqualValues(t, p.expect, r, p.desrc)
	}
}

func TestFirst(t *testing.T) {
	for _, p := range firstTestPairs {
		r := p.target.First()

		assert.EqualValues(t, p.expect, r, p.desrc)
	}
}

func TestHas(t *testing.T) {
	for _, p := range hasTestPairs {
		r := p.target.Has(p.value)

		assert.EqualValues(t, p.expect, r, p.desrc)
	}
}

func TestHasClass(t *testing.T) {
	for _, p := range hasCLassTestPairs {
		r := p.target.HasClass(p.value)

		assert.EqualValues(t, p.expect, r, p.desrc)
	}
}
