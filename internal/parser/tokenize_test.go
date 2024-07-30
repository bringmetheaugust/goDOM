package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tokenizeExpect = []string{
	"<html lang='ua'>",
	"<head>",
	"<title>",
	"Test tokenize",
	"</title>",
	"<link type='stylesheet' href='/LICENSE' />",
	"<link type='xml/text' href='/README.md' />",
	"</head>",
	"<body>",
	"<ul>",
	"<li>",
	"li 1",
	"</li>",
	"<li>",
	"li 2",
	"</li>",
	"</ul>",
	"</body>",
	"</html>",
}

func Test_tokenize(t *testing.T) {
	testFile, _ := os.ReadFile("../../test/tokenize.html")
	testFileStr := normalize(string(testFile))
	v := tokenize(testFileStr)

	assert.EqualValuesf(t, tokenizeExpect, v, "")
}
