package goDom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tagTestPair struct {
	value  string
	expect tag
}

var tagTests = [...]tagTestPair{
	{
		"<div class='lol lil lop' hidden data-set='1 2 3' data_mid='wi-fi'>",
		tag{
			name: "div",
			attributes: attributes{
				"class":    "lol lil lop",
				"hidden":   "",
				"data-set": "1 2 3",
				"data_mid": "wi-fi",
			},
			selfClosing: false,
		},
	},
	{
		"<div class='lol lil lop' hidden data-set='1 2 3' data_mid='wi-fi'/>",
		tag{
			name: "div",
			attributes: attributes{
				"class":    "lol lil lop",
				"hidden":   "",
				"data-set": "1 2 3",
				"data_mid": "wi-fi",
			},
			selfClosing: true,
		},
	},
	{
		"<div class='lol'>",
		tag{
			name:        "div",
			attributes:  attributes{"class": "lol"},
			selfClosing: false,
		},
	},
	{
		"<div class='lol'/>",
		tag{
			name:        "div",
			attributes:  attributes{"class": "lol"},
			selfClosing: true,
		},
	},
	{
		"<div>",
		tag{
			name:        "div",
			attributes:  nil,
			selfClosing: false,
		},
	},
	{
		"<div/>",
		tag{
			name:        "div",
			attributes:  nil,
			selfClosing: true,
		},
	},
}

func Test_parseTag(t *testing.T) {
	for _, pair := range tagTests {
		v := parseTag(pair.value)

		assert.EqualValuesf(t, pair.expect, v, "")
	}
}
