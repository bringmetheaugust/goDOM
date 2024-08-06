package goDom

import (
	"os"
	"testing"
)

var normalizeExpect = "<div class='lol'>lol?<span>lol!</span></div>"

func Test_normalize(t *testing.T) {
	testFile, _ := os.ReadFile("./test/normalize.html")

	if v := normalize(string(testFile)); v != normalizeExpect {
		t.Error("\nExpect", normalizeExpect, "\nGot", v)
	}
}
