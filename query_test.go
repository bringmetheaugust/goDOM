package goDom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type queryTestPair struct {
	description string
	value       string
	expect      query
	expectErr   bool
}
type queriesTestPair struct {
	description string
	value       string
	expect      queries
	expectErr   bool
}

var (
	queryTests = [...]queryTestPair{
		{
			"Tag query.",
			"div",
			query{
				tagName:   "div",
				id:        "",
				classList: nil,
				operator:  "",
				child:     nil,
			},
			false,
		},
		{
			"Class query.",
			".lol",
			query{
				tagName:   "",
				id:        "",
				classList: classList{"lol"},
				operator:  "",
				child:     nil,
			},
			false,
		},
		{
			"Id query.",
			"#lal",
			query{
				tagName:   "",
				id:        "lal",
				classList: nil,
				operator:  "",
				child:     nil,
			},
			false,
		},
		{
			"Tag with contained attribute query.",
			"div[lol]",
			query{
				tagName:    "div",
				id:         "",
				classList:  nil,
				attributes: attributes{"lol": ""},
				operator:   "",
				child:      nil,
			},
			false,
		},
		{
			"Tag with attribute query.",
			"div[lol='soccer']",
			query{
				tagName:    "div",
				id:         "",
				classList:  nil,
				attributes: attributes{"lol": "soccer"},
				operator:   "",
				child:      nil,
			},
			false,
		},
		{
			"Multi attributes.",
			"input[type='radio'][name='gender'][checked]",
			query{
				tagName:    "input",
				id:         "",
				classList:  nil,
				attributes: attributes{"type": "radio", "name": "gender", "checked": ""},
				operator:   "",
				child:      nil,
			},
			false,
		},
		{
			"Simple one stage query.",
			"ul#lal.lol-1[data='pups']",
			query{
				tagName:    "ul",
				id:         "lal",
				classList:  classList{"lol-1"},
				attributes: attributes{"data": "pups"},
				operator:   "",
				child:      nil,
			},
			false,
		},
		{
			"Complicated one stage query.",
			"ul#lal.lol-1.lol-2.lol_3[data='pups'][visible='false']",
			query{
				tagName:    "ul",
				id:         "lal",
				classList:  classList{"lol-1", "lol-2", "lol_3"},
				attributes: attributes{"data": "pups", "visible": "false"},
				operator:   "",
				child:      nil,
			},
			false,
		},
		{
			"Multistage query.",
			"ul#lal li.lol span.lol_1.lol-2",
			query{
				tagName:   "ul",
				id:        "lal",
				classList: nil,
				operator:  "",
				child: &query{
					tagName:   "li",
					id:        "",
					classList: classList{"lol"},
					operator:  "",
					child: &query{
						tagName:   "span",
						id:        "",
						classList: classList{"lol_1", "lol-2"},
						operator:  "",
						child:     nil,
					},
				},
			},
			false,
		},
		{
			"With operator.",
			"#wee * span .lol * ",
			query{
				tagName:   "",
				id:        "wee",
				classList: nil,
				operator:  "",
				child: &query{
					tagName:   "",
					id:        "",
					classList: nil,
					operator:  "*",
					child: &query{
						tagName:   "span",
						id:        "",
						classList: nil,
						operator:  "",
						child: &query{
							tagName:   "",
							id:        "",
							classList: classList{"lol"},
							operator:  "",
							child: &query{
								tagName:   "",
								id:        "",
								classList: nil,
								operator:  "*",
								child:     nil,
							},
						},
					},
				},
			},
			false,
		},
		{
			"Query shouldn't contains separated selectors.",
			"div, .lol",
			query{},
			true,
		},
		{
			"Query should pass error.",
			"",
			query{},
			true,
		},
	}
	queriesTest = []queriesTestPair{
		{
			"Paired queries with operators.",
			"div, .lol *",
			queries{
				{
					tagName:   "div",
					id:        "",
					classList: nil,
					operator:  "",
					child:     nil,
				},
				{
					tagName:   "",
					id:        "",
					classList: classList{"lol"},
					operator:  "",
					child: &query{
						tagName:   "",
						id:        "",
						classList: nil,
						operator:  "*",
						child:     nil,
					},
				},
			},
			false,
		},
		{
			"Paired queries.",
			"#hill, .billy, redneck",
			queries{
				{tagName: "", id: "hill", classList: nil, attributes: nil, child: nil},
				{tagName: "", id: "", classList: classList{"billy"}, attributes: nil, child: nil},
				{tagName: "redneck", id: "", classList: nil, attributes: nil, child: nil},
			},
			false,
		},
	}
)

func Test_parseQuery(t *testing.T) {
	for _, p := range queryTests {
		v, err := parseQuery(p.value)

		if p.expectErr {
			if err != nil {
				continue
			}

			t.Error(p.description, ": \nfor", p.value, "\nexpected error")

			continue
		}

		assert.EqualValuesf(t, &p.expect, v, p.description)
	}
}

func Test_parseQueries(t *testing.T) {
	for _, p := range queriesTest {
		v, err := parseQueries(p.value)

		if p.expectErr {
			if err != nil {
				continue
			}

			t.Error(p.description, ": \nfor", p.value, "\nexpected error")

			continue
		}

		assert.EqualValuesf(t, &p.expect, &v, p.description)
	}
}
