package goDom

import "fmt"

type JQ []*Element

// Get attribute.
// https://api.jquery.com/attr
// You cann't pass second argument to set attribute like in original jQuery cause goDOM is not modify DOM.
//
// Examples:
//
//	_, jQ, _ := goDom.Create(bytes)
//	attr, err := jQ(".my-class").Attr("scr") // err if no attribute
func (j *JQ) Attr(attr string) (string, error) {
	if len(*j) == 0 {
		return "", notFoundErr{Msg: "elements by query not found"}
	}

	return (*j)[0].GetAttribute(attr)
}

// Get children from each elements
// https://api.jquery.com/children
//
// Examples:
//
//	// get children from each jQ elements
//	_, jQ, _ := goDom.Create(bytes)
//	c, err := jQ(".my-class").Children() // err if no children
//
//	// get children from each jQ elements which has class ".lol"
//
//	_, jQ, _ := goDom.Create(bytes)
//	c, err := jQ(".my-class").Children(".lol") // err if no children
func (j *JQ) Children(a ...string) JQ {
	var jq JQ

	for _, v := range *j {
		if len(v.Children) > 0 {
			if len(a) == 0 {
				jq = append(jq, v.Children...)
			} else {
				// TODO: scope & >
				if r, err := v.QuerySelectorAll(fmt.Sprintf(":scope > %v", a[0])); err == nil {
					jq = append(jq, r...)
				}
			}
		}
	}

	return jq
}

// Loop all jQuery elements with callback.
// https://api.jquery.com/each
//
// Examples:
//
//	_, jQ, _ := goDom.Create(bytes)
//	jQ("div").Each(func(i *goDom.JQ) {
//		attr, _ := i.Attr("class")
//		print(attr)
//	})
func (j *JQ) Each(f func(i *JQ)) {
	for _, v := range *j {
		f(JQ{}.new(v))
	}
}

// Filter elements.
// https://api.jquery.com/filter
//
// Examples:
//
//	// filter element which correspond/match query selector
//	_, jQ, _ := goDom.Create(bytes)
//	jQ("div").Filter(".oops") // filter elements with ".oops" classes
//
//	// filter elements by callback function
//	_, jQ, _ := goDom.Create(bytes)
//	jQ("div").Filter(func(i goDom.JQ) bool {
//		_, err := i.Attr("class")
//		return err != nil // filter elements which have class attribute
//	})
func (j *JQ) Filter(i interface{}) JQ {
	var jq JQ

	switch v := i.(type) {
	case string:
		for _, e := range *j {
			// TODO multi queries
			if elementMatchesQuery(*createQuery(v), e) {
				jq = append(jq, e)
			}
		}
	case func(i *JQ) bool:
		for _, e := range *j {
			if v(JQ{}.new(e)) {
				jq = append(jq, e)
			}
		}
	default:
		return nil
	}

	return jq
}

// Find elements inside each jQuery element.
// https://api.jquery.com/find
//
// Example:
//
//	_, jQ, _ := goDom.Create(bytes)
//	JQ(".piu").Find("li.wee") // inside each element with ".piu" classes find all "li" elements with "wee" classes
func (j *JQ) Find(q string) JQ {
	var jq JQ

	for _, v := range *j {
		// TODO multi queries
		if r, err := v.QuerySelectorAll(q); err == nil {
			jq = append(jq, r...)
		}
	}

	return jq
}

// Get first element from jQuery elements.
// https://api.jquery.com/first
//
// Examples:
//
//	_, jQ, _ := goDom.Create(bytes)
//	jQ("li").First()
func (j *JQ) First() JQ {
	if len((*j)) > 0 {
		return *JQ{}.new((*j)[0])
	}

	return *JQ{}.new()
}

// Filter jQuery elements which have another elements inside children.
// https://api.jquery.com/has
//
// Examples:
//
//	// has elemenst by query selector
//	_, jQ, _ := goDom.Create(bytes)
//	result := jQ("li").has(".boom"); // get all `li` elements which contains inside itself `.boom` elements
//
//	// contains another jQuery inside itself
//	_, jQ, _ := goDom.Create(bytes)
//	dick := jQ(".dick")
//	result := jQ(".man").Has(dick) // get all `.man` elements which contains inside itself `.dick` jQuery
func (j *JQ) Has(i interface{}) JQ {
	var jQ JQ

	switch v := i.(type) {
	case string:
		for _, e := range *j {
			if _, err := e.QuerySelector(v); err == nil {
				jQ = append(jQ, e)
			}
		}
	case *JQ:
		for _, root := range *j {
			for _, nested := range *v {
				if root.Contains(nested) {
					jQ = append(jQ, root)
				}
			}
		}
	default:
		jQ = *JQ{}.new()
	}

	return jQ
}

func (j JQ) new(e ...*Element) *JQ {
	var jq JQ

	jq = append(jq, e...)

	return &jq
}
