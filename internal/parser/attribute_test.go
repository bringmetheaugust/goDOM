package parser

import (
	"goDOM/internal/dom"
	"testing"

	"github.com/stretchr/testify/assert"
)

type attrTestPair struct {
	value  string
	expect dom.Attribute
}

var attributeTests = []attrTestPair{
	{"class='lol'", dom.Attribute{Name: "class", Value: "lol"}},
	{"href='htttp://lol.com'", dom.Attribute{Name: "href", Value: "htttp://lol.com"}},
	{"data-name='tag id'", dom.Attribute{Name: "data-name", Value: "tag id"}},
	{"data_shit='bitch'", dom.Attribute{Name: "data_shit", Value: "bitch"}},
	{"hidden", dom.Attribute{Name: "hidden", Value: ""}},
	{"hidden=''", dom.Attribute{Name: "hidden", Value: ""}},
	{"hidden='? __ ? &__'", dom.Attribute{Name: "hidden", Value: "? __ ? &__"}},
}

func Test_parseAttribute(t *testing.T) {
	for _, pair := range attributeTests {
		v := parseAttribute(pair.value)

		assert.EqualValuesf(t, pair.expect, v, "")
	}
}
