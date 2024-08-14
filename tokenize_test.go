package goDom

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tokenizeExpect = []token{
	{token_element, "<!DOCTYPE html>"},
	{token_element, "<html lang='ua'>"},
	{token_element, "<head>"},
	{token_element, "<title>"},
	{token_element, "Test tokenize"},
	{token_element, "</title>"},
	{token_element, "<link type='stylesheet' href='/LICENSE' />"},
	{token_element, "<link type='xml/text' href='/README.md' />"},
	{token_element, "<style id='wp-block-library-theme-inline-css'>"},
	{token_content, "a > div { display: none !important; } .wp-block-search__button { border: 1px solid #ccc; padding: .375em .625em } :where(.wp-block-group.has-background) { padding: 1.25em 2.375em } .wp-block-separator.has-background:not(.is-style-wide):not(.is-style-dots) { height: 2px }"},
	{token_element, "</style>"},
	{token_element, "<style id='style_with_attribute'>"},
	{token_content, ".lol { display: flex; }"},
	{token_element, "</style>"},
	{token_element, "</head>"},
	{token_element, "<body>"},
	{token_element, "<!-- <div>wtf, man???</div> -->"},
	{token_element, "<ul>"},
	{token_element, "<li>"},
	{token_element, "li 1"},
	{token_element, "</li>"},
	{token_element, "<li>"},
	{token_element, "li 2"},
	{token_element, "<!-- <ul><li></li><li>блять..</li></ul> -->"},
	{token_element, "</li>"},
	{token_element, "</ul>"},
	{token_element, "</body>"},
	{token_element, "<script>"},
	{token_content, "console.log(2 > 1 != 1 --)"},
	{token_element, "</script>"},
	{token_element, "<script id='script_with_attribute'>"},
	{token_content, "console.log(1 < 2)"},
	{token_element, "</script>"},
	{token_element, "</html>"},
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
