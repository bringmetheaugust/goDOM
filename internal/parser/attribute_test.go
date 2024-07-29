package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type attrTestPair struct {
	value  string
	expect attribute
}

var attributeTests = [...]attrTestPair{
	{"class='lol'", attribute{name: "class", value: "lol"}},
	{"href='htttp://lol.com'", attribute{name: "href", value: "htttp://lol.com"}},
	{"data-name='tag id'", attribute{name: "data-name", value: "tag id"}},
	{"data_shit='bitch'", attribute{name: "data_shit", value: "bitch"}},
	{"hidden", attribute{name: "hidden", value: ""}},
	{"hidden=''", attribute{name: "hidden", value: ""}},
	{"hidden='? __ ? &__'", attribute{name: "hidden", value: "? __ ? &__"}},
}

func Test_parseAttribute(t *testing.T) {
	for _, pair := range attributeTests {
		v := parseAttribute(pair.value)

		assert.EqualValuesf(t, pair.expect, v, "")
	}
}
