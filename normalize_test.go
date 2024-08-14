package goDom

import (
	"os"
	"strings"
	"testing"
)

var normalizeExpect = "<div class='lol'>lol?<span>lol!</span></div>"

func Test_normalize(t *testing.T) {
	testFile, _ := os.ReadFile("./test/normalize.html")
	ch := make(chan string)
	var res strings.Builder

	go normalize(string(testFile), ch)

	for r := range ch {
		res.WriteString(r)
	}

	if res.String() != normalizeExpect {
		t.Error("\nExpect", normalizeExpect, "\nGot", res)
	}
}
