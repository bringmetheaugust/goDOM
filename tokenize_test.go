package goDom

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tokenizeExpect = []token{
	{_type: node_meta, data: "<!DOCTYPE html>"},
	{_type: node_element, tag: parseTag("<html lang='ua'>")},
	{_type: node_element, tag: parseTag("<head>")},
	{_type: node_element, tag: parseTag("<title>")},
	{_type: node_text, data: "Test tokenize"},
	{_type: node_element, isClosing: true}, // </title>
	{_type: node_element, tag: parseTag("<link type='stylesheet' href='/LICENSE' />")},
	{_type: node_element, tag: parseTag("<link type='xml/text' href='/README.md' />")},
	{_type: node_element, tag: parseTag("<style id='wp-block-library-theme-inline-css'>")},
	{_type: node_text, data: "a > div { display: none !important; } .wp-block-search__button { border: 1px solid #ccc; padding: .375em .625em } :where(.wp-block-group.has-background) { padding: 1.25em 2.375em } .wp-block-separator.has-background:not(.is-style-wide):not(.is-style-dots) { height: 2px }"},
	{_type: node_element, isClosing: true}, // </style>
	{_type: node_element, tag: parseTag("<style id='style_with_attribute'>")},
	{_type: node_text, data: ".lol { display: flex; }"},
	{_type: node_element, isClosing: true}, // </style>
	{_type: node_element, isClosing: true}, // </head>
	{_type: node_element, tag: parseTag("<body>")},
	{_type: node_comment, data: "<!-- <div>wtf, man???</div> -->"},
	{_type: node_element, tag: parseTag("<ul>")},
	{_type: node_element, tag: parseTag("<li>")},
	{_type: node_text, data: "li 1"},
	{_type: node_element, isClosing: true}, // </li>
	{_type: node_element, tag: parseTag("<li>")},
	{_type: node_text, data: "li 2"},
	{_type: node_comment, data: "<!-- <ul><li></li><li>блять..</li></ul> -->"},
	{_type: node_element, isClosing: true}, // </li>
	{_type: node_element, isClosing: true}, // </ul>
	{_type: node_element, isClosing: true}, // </body>
	{_type: node_element, tag: parseTag("<script>")},
	{_type: node_text, data: "console.log(2 > 1 != 1 --)"},
	{_type: node_element, isClosing: true}, // </script>
	{_type: node_element, tag: parseTag("<script id='script_with_attribute'>")},
	{_type: node_text, data: "console.log(1 < 2)"},
	{_type: node_element, isClosing: true}, // </script>
	{_type: node_element, isClosing: true}, // </html>
}

func Test_tokenize(t *testing.T) {
	testFile, _ := os.ReadFile("./test/tokenize.html")
	chMarkupLine := make(chan string)
	chTokens := make(chan token)
	var tokens []token

	go normalize(string(testFile), chMarkupLine)
	go tokenize(chMarkupLine, chTokens)

	for token := range chTokens {
		tokens = append(tokens, token)
	}

	assert.EqualValuesf(t, tokenizeExpect, tokens, "")
}
