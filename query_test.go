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

var queryTests = [...]queryTestPair{
	{
		"Tag query.",
		"div",
		query{tagName: "div", id: "", classList: nil, child: nil},
		false,
	},
	{
		"Class query.",
		".lol",
		query{tagName: "", id: "", classList: classList{"lol"}, child: nil},
		false,
	},
	{
		"Id query.",
		"#lal",
		query{tagName: "", id: "lal", classList: nil, child: nil},
		false,
	},
	{
		"Tag with contained attribute query.",
		"div[lol]",
		query{tagName: "div", id: "", classList: nil, attributes: attributes{"lol": ""}, child: nil},
		false,
	},
	{
		"Tag with attribute query.",
		"div[lol='soccer']",
		query{tagName: "div", id: "", classList: nil, attributes: attributes{"lol": "soccer"}, child: nil},
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
			child: &query{
				tagName:   "li",
				id:        "",
				classList: classList{"lol"},
				child: &query{
					tagName:   "span",
					id:        "",
					classList: classList{"lol_1", "lol-2"},
					child:     nil,
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
var queriesTest = []queriesTestPair{
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

func Test_parseQuery(t *testing.T) {
	for _, pair := range queryTests {
		v, err := parseQuery(pair.value)

		if pair.expectErr {
			if err != nil {
				continue
			}

			t.Error(pair.description, ": \nfor", pair.value, "\nexpected error")

			continue
		}

		assert.EqualValuesf(t, &pair.expect, v, pair.description)
	}
}

func Test_parseQueries(t *testing.T) {
	for _, pair := range queriesTest {
		v, err := parseQueries(pair.value)

		if pair.expectErr {
			if err != nil {
				continue
			}

			t.Error(pair.description, ": \nfor", pair.value, "\nexpected error")

			continue
		}

		assert.EqualValuesf(t, &pair.expect, &v, pair.description)
	}
}
