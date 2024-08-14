package goDom

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tokenizeExpect = []string{
	"<!DOCTYPE html>",
	"<html lang='ua'>",
	"<head>",
	"<title>",
	"Test tokenize",
	"</title>",
	"<link type='stylesheet' href='/LICENSE' />",
	"<link type='xml/text' href='/README.md' />",
	"<style id='wp-block-library-theme-inline-css'>",
	"a > div {display: none !important;}.wp-block-search__button {border: 1px solid #ccc;padding: .375em .625em}:where(.wp-block-group.has-background) {padding: 1.25em 2.375em}.wp-block-separator.has-background:not(.is-style-wide):not(.is-style-dots) {height: 2px}",
	"</style>",
	"<style id='style_with_attribute'>",
	".lol {display: flex;}",
	"</style>",
	"</head>",
	"<body>",
	"<!-- <div>wtf, man???</div> -->",
	"<ul>",
	"<li>",
	"li 1",
	"</li>",
	"<li>",
	"li 2",
	"<!-- <ul><li></li><li>блять..</li></ul> -->",
	"</li>",
	"</ul>",
	"</body>",
	"<script>",
	"console.log(2 > 1 != 1 --)",
	"</script>",
	"<script id='script_with_attribute'>",
	"console.log(1 < 2)",
	"</script>",
	"</html>",
}

func Test_tokenize(t *testing.T) {
	testFile, _ := os.ReadFile("./test/tokenize.html")
	chMarkupLine := make(chan string)
	chTokens := make(chan string)
	var tokens []string

	go normalize(string(testFile), chMarkupLine)
	go tokenize(chMarkupLine, chTokens)

	for token := range chTokens {
		tokens = append(tokens, token)
	}

	assert.EqualValuesf(t, tokenizeExpect, tokens, "")
}
