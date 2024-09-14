package goDom

import (
	"slices"
	"strings"
)

type JQ []*Element

// Get attribute from first element in jQuery elements.
// https://api.jquery.com/attr
//
// Examples:
//
//	attr, err := jQ(".my-life").Attr("scr") // err if no attribute
func (j *JQ) Attr(attr string) (string, error) {
	if len(*j) == 0 {
		return "", notFoundErr{Msg: "elements by query not found"}
	}

	return (*j)[0].GetAttribute(attr)
}

// Get children from each jQuery elements. It also support multi queries (".my-class, div.boom").
// https://api.jquery.com/children
//
// Examples:
//
//	// get all childrens
//	c := jQ(".my-boobs").Children()
//
//	// get children using query
//	c := jQ(".my-class").Children(".lol")
//	c := jQ(".my-class, div.boom").Children(".lol") // using multi queries
func (j *JQ) Children(a ...string) JQ {
	var jq JQ

	if len(a) == 0 {
		for _, v := range *j {
			jq = append(jq, v.Children...)
		}
	} else {
		if queries, err := parseQueries(a[0]); err == nil {
			for _, e := range *j {
				for _, q := range queries {
					for _, c := range e.Children {
						if c.elementMatchesQuery(q, c) {
							jq = append(jq, c)
						}
					}
				}
			}
		}
	}

	return jq
}

// Loop all jQuery elements with callback function.
// https://api.jquery.com/each
//
// Examples:
//
//	jQ("div").Each(func(i *goDom.JQ) {
//		attr, _ := i.Attr("class")
//		print(attr)
//	})
func (j *JQ) Each(f func(i *JQ)) {
	for _, v := range *j {
		f(JQ{}.new(v))
	}
}

// Filter jQuery elements which correspond/match query selector or callback function.
// It also supports multi queries (".oops, li.pow").
// https://api.jquery.com/filter
//
// Examples:
//
//	// using query
//	jQ("div").Filter(".oops")
//	jQ("div").Filter(".oops, li.dow") // using multi queries
//
//	// using callback function
//	jQ("div").Filter(func(i goDom.JQ) bool {
//		_, err := i.Attr("class")
//		return err != nil
//	})
func (j *JQ) Filter(i interface{}) JQ {
	var jq JQ

	switch v := i.(type) {
	case string:
		for _, e := range *j {
			if queries, err := parseQueries(v); err == nil {
				for _, q := range queries {
					if e.elementMatchesQuery(q, e) {
						jq = append(jq, e)
					}
				}
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

// Find elements inside each jQuery element. Also support multi queries("li.lal, #lol")
// https://api.jquery.com/find
//
// Example:
//
//	JQ(".piu").Find("li.wee")
//	JQ(".piu").Find("li.wee, #pow") // using multi queries
func (j *JQ) Find(query string) (jq JQ) {
	for _, e := range *j {
		splitedQueries := strings.Split(query, ",")

		for _, q := range splitedQueries {
			if r, err := e.QuerySelectorAll(q); err == nil {
				jq = append(jq, r...)
			}
		}
	}

	return
}

// Get first element from jQuery elements.
// https://api.jquery.com/first
//
// Examples:
//
//	jQ("li").First()
func (j *JQ) First() JQ {
	if len((*j)) > 0 {
		return *JQ{}.new((*j)[0])
	}

	return *JQ{}.new()
}

// Filter jQuery elements which have another elements inside their children.
// https://api.jquery.com/has
//
// Examples:
//
//	// using query
//	result := jQ("li").has(".boom");
//
//	// using jQuery element
//	dick := jQ(".dick")
//	result := jQ(".man").Has(dick)
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

// Check if element contains class. Return true is contains.
// https://api.jquery.com/hasClass
//
// Examples:
//
//	res := jQ("#cat").HasClass("meow")
func (j *JQ) HasClass(class string) bool {
	if len((*j)) > 0 {
		return slices.Contains((*j)[0].ClassList, class)
	}

	return false
}

func (j JQ) new(e ...*Element) *JQ {
	var jq JQ

	jq = append(jq, e...)

	return &jq
}
