package goDom

import "testing"

type elementTestPair[E string | bool] struct {
	value     string
	expect    E
	expectErr bool
}

var getAttributeTests = []elementTestPair[string]{
	{"href", "http://pizdets.com", false},
	{"hidden", "", false},
	{"magic", "", true},
	{"", "", true},
}
var hasAttributeTests = []elementTestPair[bool]{
	{value: "href", expect: true},
	{value: "hidden", expect: true},
	{value: "magic", expect: false},
	{value: "", expect: false},
}
var mockElement = Element{
	Attributes: attributes{
		"href":   "http://pizdets.com",
		"hidden": "",
	},
}

func TestGetAttribute(t *testing.T) {
	for _, pair := range getAttributeTests {
		v, err := mockElement.GetAttribute(pair.value)

		if pair.expectErr {
			if err != nil {
				continue
			}

			t.Error("\nfor", pair.value, "\nexpected error")
		} else {
			if v != pair.expect {
				t.Error("\nExpect", pair.expect, "\nGot", v)
			}
		}
	}
}

func TestHasAttribute(t *testing.T) {
	for _, pair := range hasAttributeTests {
		if v := mockElement.HasAttribute(pair.value); v != pair.expect {
			t.Error("\nExpect", pair.expect, "\nGot", v)
		}
	}
}
