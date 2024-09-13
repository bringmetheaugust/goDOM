package goDom

import "testing"

type elementTestPair[E string | bool] struct {
	value     string
	expect    E
	expectErr bool
}

var (
	getAttributeTests = []elementTestPair[string]{
		{value: "src", expect: "http://zalupa.img.com"},
		{value: "hidden", expect: ""},
		{"magic", "", true},
		{"", "", true},
	}
	hasAttributeTests = []elementTestPair[bool]{
		{value: "data-custom", expect: true},
		{value: "disabled", expect: true},
		{value: "magic", expect: false},
		{value: "", expect: false},
	}
)

func TestGetAttribute(t *testing.T) {
	for _, p := range getAttributeTests {
		v, err := mockEl_img_1.GetAttribute(p.value)

		if p.expectErr {
			if err != nil {
				continue
			}

			t.Error("\nfor", p.value, "\nexpected error")
		} else {
			if v != p.expect {
				t.Error("\nExpect", p.expect, "\nGot", v)
			}
		}
	}
}

func TestHasAttribute(t *testing.T) {
	for _, p := range hasAttributeTests {
		if v := mockEl_button_1.HasAttribute(p.value); v != p.expect {
			t.Error("\nExpect", p.expect, "\nGot", v)
		}
	}
}
