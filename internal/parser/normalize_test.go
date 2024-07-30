package parser

import "testing"

var normalizeTest = `   <div class="lol"></div>  `
var normalizeExpect = "<div class='lol'></div>"

func Test_normalize(t *testing.T) {
	if v := normalize(normalizeTest); v != normalizeExpect {
		t.Error(
			"\nExpect",
			normalizeExpect,
			"\nGot",
			v,
		)
	}
}
