package parser

import (
	"goDOM/internal/dom"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tagTestPair struct {
	value  string
	expect tag
}

var tagTests = [...]tagTestPair{
	{
		"div class='lol lil lop' hidden data-set='1 2 3' data_mid='wi-fi'",
		tag{name: "div", attributes: dom.Attributes{
			"class":    "lol lil lop",
			"hidden":   "",
			"data-set": "1 2 3",
			"data_mid": "wi-fi",
		}},
	},
	{
		"div class='lol'",
		tag{name: "div", attributes: dom.Attributes{
			"class": "lol",
		}},
	},
	{
		"div",
		tag{"div", dom.Attributes{}},
	},
}

func Test_parseTag(t *testing.T) {
	for _, pair := range tagTests {
		v := parseTag(pair.value)

		assert.EqualValuesf(t, pair.expect, v, "")
	}
}
