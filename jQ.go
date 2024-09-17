package goDom

import (
	"slices"
	"strings"
)

type JQ []*Element

// Get the value of an attribute for the first element in the set of matched elements
// or set one or more attributes for every matched element.
//
// https://api.jquery.com/attr
//
// Examples:
//
//	attr, err := jQ(".my-life").Attr("scr") // err if no attribute
func (j JQ) Attr(attr string) (string, error) {
	if len(j) == 0 {
		return "", notFoundErr{Msg: "elements by query not found"}
	}

	return j[0].GetAttribute(attr)
}

// Get the children of each element in the set of matched elements, optionally filtered by a selector.
// It also support multi selectors (".my-class, div.boom").
// https://api.jquery.com/children
//
// Examples:
//
//	// get all childrens
//	c := jQ(".my-boobs").Children()
//
//	// get children using selector
//	c := jQ(".my-class").Children(".lol")
//	c := jQ(".my-class, div.boom").Children(".lol") // or using multi selectors
func (j JQ) Children(selectors ...string) JQ {
	var jq JQ

	if len(selectors) == 0 {
		for _, v := range j {
			jq = append(jq, v.Children...)
		}
	} else {
		if queries, err := parseQueries(selectors[0]); err == nil {
			for _, e := range j {
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

// Iterate over a jQuery object, executing a function for each matched element.
// https://api.jquery.com/each
//
// Examples:
//
//	jQ("div").Each(func(i *goDom.JQ) {
//		attr, _ := i.Attr("class")
//		print(attr)
//	})
func (j JQ) Each(f func(i JQ)) {
	for _, v := range j {
		f(JQ{}.new(v))
	}
}

// Reduce the set of matched elements to those that match the selector or pass the function's test.
// It also supports multi selectors (".oops, li.pow").
// https://api.jquery.com/filter
//
// Examples:
//
//	// using selector
//	jQ("div").Filter(".oops")
//	jQ("div").Filter(".oops, li.dow") // or using multi selectors
//
//	// using callback function
//	jQ("div").Filter(func(i goDom.JQ) bool {
//		_, err := i.Attr("class")
//		return err != nil
//	})
func (j JQ) Filter(i interface{}) JQ {
	var jq JQ

	switch v := i.(type) {
	case string:
		for _, e := range j {
			if queries, err := parseQueries(v); err == nil {
				for _, q := range queries {
					if e.elementMatchesQuery(q, e) {
						jq = append(jq, e)
					}
				}
			}
		}
	case func(i JQ) bool:
		for _, e := range j {
			if v(JQ{}.new(e)) {
				jq = append(jq, e)
			}
		}
	default:
		return nil
	}

	return jq
}

// Get the descendants of each element in the current set of matched elements,
// filtered by a selector. Also support multi selectors ("li.lal, #lol")
// https://api.jquery.com/find
//
// Example:
//
//	JQ(".piu").Find("li.wee")
//	JQ(".piu").Find("li.wee, #pow") // using multi selectors
func (j JQ) Find(selectors string) (jq JQ) {
	for _, e := range j {
		splitedQueries := strings.Split(selectors, ",")

		for _, q := range splitedQueries {
			if r, err := e.QuerySelectorAll(q); err == nil {
				jq = append(jq, r...)
			}
		}
	}

	return
}

// Reduce the set of matched elements to the first in the set.
// https://api.jquery.com/first
//
// Examples:
//
//	jQ("li").First()
func (j JQ) First() JQ {
	if len(j) > 0 {
		return JQ{}.new(j[0])
	}

	return JQ{}.new()
}

// Reduce the set of matched elements to those that have a descendant that matches the selector or DOM element.
// https://api.jquery.com/has
//
// Examples:
//
//	// using selector
//	result := jQ("li").has(".boom");
//
//	// using jQuery element
//	dick := jQ(".dick")
//	result := jQ(".man").Has(dick)
func (j JQ) Has(i interface{}) JQ {
	var jQ JQ

	switch v := i.(type) {
	case string:
		for _, e := range j {
			if _, err := e.QuerySelector(v); err == nil {
				jQ = append(jQ, e)
			}
		}
	case JQ:
		for _, root := range j {
			for _, nested := range v {
				if root.Contains(nested) {
					jQ = append(jQ, root)
				}
			}
		}
	default:
		jQ = JQ{}.new()
	}

	return jQ
}

// Determine whether any of the matched elements are assigned the given class.
// https://api.jquery.com/hasClass
//
// Examples:
//
//	res := jQ("#cat").HasClass("meow")
func (j JQ) HasClass(class string) bool {
	if len(j) > 0 {
		return slices.Contains(j[0].ClassList, class)
	}

	return false
}

// Reduce the set of matched elements to the final one in the set.
// https://api.jquery.com/last
//
// Examples:
//
//	res := jQ(".lol").Last()
func (j JQ) Last() JQ {
	var jQ JQ

	if len(j) != 0 {
		jQ = append(jQ, j[len(j)-1])
	}

	return jQ
}

// Get the immediately following sibling of each element in the set of matched elements.
// If a selector is provided, it retrieves the next sibling only if it matches that selector.
// https://api.jquery.com/next
//
// Examples:
//
//	// using without selector
//	res := jQ("li").Next()
//
//	// using selector
//	res := jQ("li").Next(".lol")
func (j JQ) Next(selectors ...string) JQ {
	return j.next(true, selectors...)
}

// Get all following siblings of each element in the set of matched elements, optionally filtered by a selector.
// https://api.jquery.com/nextAll
//
// Examples:
//
//	// using without selector
//	res := jQ("li").NextAll()
//
//	// using selector
//	res := jQ("li").NextAll(".lol")
func (j JQ) NextAll(selectors ...string) JQ {
	return j.next(false, selectors...)
}

func (j JQ) next(first bool, selectors ...string) JQ {
	var jQ JQ

	switch {
	case len(j) == 0:
		break
	case len(selectors) == 0:
		for _, e := range j {
			for s := e.NextElementSibling; ; {
				if s != nil {
					jQ = append(jQ, s)

					if first {
						break
					}

					s = s.NextElementSibling
				} else {
					break
				}
			}
		}
	case len(selectors) > 0:
		if q, err := parseQuery(selectors[0]); err == nil {
			for _, e := range j {
				for s := e.NextElementSibling; ; {
					if s != nil {
						if (domAPIUtils{}.elementMatchesQuery(*q, s)) {
							jQ = append(jQ, s)

							if first {
								break
							}
						}

						s = s.NextElementSibling
					} else {
						break
					}
				}
			}
		}
	}

	return jQ
}

// Remove elements from the set of matched elements.
// https://api.jquery.com/not
//
// Examples:
//
//	// using selector
//	res := jQ("li").Not(".bird")
//	res := jQ("li").Not(".bird, .bear") // or using multi selectors
//
//	// using callback function
//	res := jQ("li").Not(func(el JQ) bool {
//		return el.HasClass("bird")
//	})
//
//	// using jQuery element
//	birds := jQ(".bird")
//	res := jQ("li").Not(birds)
func (j JQ) Not(i interface{}) JQ {
	var jq JQ

	if len(j) != 0 {
		switch v := i.(type) {
		case string:
			if queries, err := parseQueries(v); err == nil {
			eLoop:
				for _, e := range j {
					for _, q := range queries {
						if (domAPIUtils{}.elementMatchesQuery(q, e)) {
							continue eLoop
						}
					}

					jq = append(jq, e)
				}
			}
		case func(el JQ) bool:
			for _, e := range j {
				if !v(JQ{}.new(e)) {
					jq = append(jq, e)
				}
			}
		case JQ:
			for _, e := range j {
				if !slices.Contains(v, e) {
					jq = append(jq, e)
				}
			}
		}
	}

	return jq
}

// Get the parent of each element in the current set of matched elements, optionally filtered by a selector.
// https://api.jquery.com/parent
//
// Examples:
//
//	// using without selectors
//	res := jQ("li").Parent()
//
//	// using selector
//	res := jQ("li").Parent(".wow")
//	res := jQ("li").Parent(".wow, .wee") // or using multi selectors
func (j JQ) Parent(selector ...string) JQ {
	var jq JQ

	if len(j) != 0 {
		if len(selector) == 0 {
			for _, e := range j {
				if p := e.ParentElement; p != nil {
					jq = append(jq, p)
				}
			}
		} else {
			if queries, err := parseQueries(selector[0]); err == nil {
				for _, e := range j {
					if p := e.ParentElement; p != nil {
						for _, q := range queries {
							if (domAPIUtils{}.elementMatchesQuery(q, p)) {
								jq = append(jq, p)
							}
						}
					}
				}
			}
		}
	}

	return jq
}

// Get the ancestors of each element in the current set of matched elements, optionally filtered by a selector.
// https://api.jquery.com/parents
//
// Examples:
//
//	// using without selectors
//	res := jQ("#me").Parents()
//
//	// using selector
//	res := jQ("#me").Parents(".lol")
//	res := jQ("#me").Parents(".lol, .lil") // or using multi selectors
func (j JQ) Parents(selector ...string) JQ {
	var jq JQ

	if len(j) != 0 {
		if len(selector) == 0 {
			for _, e := range j {
				for p := e.ParentElement; ; {
					if p != nil {
						if !slices.Contains(jq, p) {
							jq = append(jq, p)
						}

						p = p.ParentElement
					} else {
						break
					}
				}
			}
		} else {
			if queries, err := parseQueries(selector[0]); err == nil {
				for _, e := range j {
					for p := e.ParentElement; ; {
						if p != nil {
							for _, q := range queries {
								if (domAPIUtils{}.elementMatchesQuery(q, p)) {
									jq = append(jq, p)
								}
							}
						} else {
							break
						}

						p = p.ParentElement
					}
				}
			}
		}
	}

	return jq
}

// Return text node from jQuery first element.
//
// Examples:
//
//	res, err := jQ("#io").Text() // error if no elements in jQuery
func (j JQ) Text() (string, error) {
	if len(j) == 0 {
		return "", notFoundErr{Msg: "No elements in jQuery."}
	}

	return j[0].TextContent, nil
}

func (j JQ) new(e ...*Element) JQ {
	var jq JQ

	jq = append(jq, e...)

	return jq
}
