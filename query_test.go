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

var queryTests = [...]queryTestPair{
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
				classList: []string{"lol"},
				child: &query{
					tagName:   "span",
					id:        "",
					classList: []string{"lol_1", "lol-2"},
					child:     nil,
				},
			},
		},
		false,
	},
	{
		"One stage query.",
		"ul#lal.lol-1.lol-2.lol_3",
		query{tagName: "ul", id: "lal", classList: []string{"lol-1", "lol-2", "lol_3"}, child: nil},
		false,
	},
	{
		"Compltex query.",
		"ul#lal.lol",
		query{tagName: "ul", id: "lal", classList: []string{"lol"}, child: nil},
		false,
	},
	{
		"Tag query.",
		"div",
		query{tagName: "div", id: "", classList: nil, child: nil},
		false,
	},
	{
		"Class query.",
		".lol",
		query{tagName: "", id: "", classList: []string{"lol"}, child: nil},
		false,
	},
	{
		"Id query.",
		"#lal",
		query{tagName: "", id: "lal", classList: nil, child: nil},
		false,
	},
	{
		"Query should pass error.",
		"",
		query{},
		true,
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
