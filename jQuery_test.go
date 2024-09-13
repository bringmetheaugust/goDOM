package goDom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type jQueryTestPair struct {
	value  any
	expect *JQ
	descr  string
}

var (
	jQueryTestElement, _  = mockDOM.QuerySelector(".top")
	jQueryTestElements, _ = mockDOM.QuerySelectorAll(".bee")
	jQueryTestPairs       = []jQueryTestPair{
		{
			value:  ".bee",
			expect: JQ{}.new(jQueryTestElements...),
			descr:  "Create jQ with search query.",
		},
		{
			value:  jQueryTestElement,
			expect: JQ{}.new(jQueryTestElement),
			descr:  "Create jQ with existed DOM element.",
		},
		{
			value:  ".nonexisted_query",
			expect: &JQ{},
			descr:  "Empty JQ.",
		},
		{
			value:  5,
			expect: &JQ{},
			descr:  "Empty JQ.",
		},
	}
)

func Test_createJQuery(t *testing.T) {
	for _, p := range jQueryTestPairs {
		jq := createJQuery(&mockDOM)
		v := jq(p.value)

		assert.EqualValuesf(t, p.expect, v, p.descr)
	}
}
