package goDom

import (
	"regexp"
	"strings"
)

type query struct {
	tagName string
	id      string
	classList
	attributes
	child *query // next level query (ex in ".lol .lal" .lal is a next level query)
}
type queries []query

// Parse and get slice of queries.
// This need in case then query has multiple separated selectors (ex for QuerySelectorAll(".lol, #hell, div")).
// Return slice which contains queries for each selector.
func parseQueries(qStr string) (queries, error) {
	if qStr == "" {
		return nil, invalidQueryErr{QueryStr: qStr, Msg: "Query is empty."}
	}

	var queries queries
	splitedQueries := strings.Split(qStr, ",")

	for _, q := range splitedQueries {
		qq, err := parseQuery(q)

		if err != nil {
			continue
		}

		queries = append(queries, *qq)
	}

	return queries, nil
}

// Parse query
func parseQuery(qStr string) (*query, error) {
	if qStr == "" {
		return &query{}, invalidQueryErr{QueryStr: qStr, Msg: "Query is empty."}
	}

	// check if query contains "," (for QuerySelectorAll)
	if splitedQueries := strings.Split(qStr, ","); len(splitedQueries) > 1 {
		return &query{}, invalidQueryErr{QueryStr: qStr, Msg: "Query contains separation for selectors. It's only possible for QuerySelectorAll."}
	}

	queries := strings.Fields(qStr)

	return createQuery(queries...), nil
}

// Create query struct
func createQuery(qArr ...string) *query {
	if len(qArr) == 0 {
		return nil
	}

	q := qArr[0]
	var newQ query

	// * get tagName

	matchedTag := matchedQueryParam(`^(\w+)(?:\.|\[|#|$)`, q)

	if len(matchedTag) > 0 {
		newQ.tagName = matchedTag[0][1]
	}

	// * get classes

	var classList classList
	matchedClasses := matchedQueryParam(`(?:\.)([\w-]+)`, q)

	for _, class := range matchedClasses {
		classList = append(classList, class[1])
	}

	newQ.classList = classList

	// * get id

	matchedId := matchedQueryParam(`(?:#)([\w-]+)`, q)

	if len(matchedId) > 0 {
		newQ.id = matchedId[0][1]
	}

	newQ.child = createQuery(qArr[1:]...)

	// * get attributes

	matchedAttrs := matchedQueryParam(`\[(\w+=[\'"][^\'"]*[\'"]|[\w-]+)\]`, q)

	if len(matchedAttrs) > 0 {
		newQ.attributes = attributes{}

		for _, m := range matchedAttrs {
			newAttr := parseAttribute(m[1])
			newQ.attributes[newAttr.name] = newAttr.value
		}
	}

	return &newQ
}

// Get query param from regExp
func matchedQueryParam(str string, q string) [][]string {
	reTag := regexp.MustCompile(str)

	return reTag.FindAllStringSubmatch(q, -1)
}

// Create query struct.
// With version of createQuery is not relevant cause have hard regExp and makes a lot of problems.
// Deprecated.
// func createQuery(qArr ...string) *query {
// 	if len(qArr) == 0 {
// 		return nil
// 	}

// 	var newQ query
// 	a := matchedQueryParam(`^(?P<Tag>\w+)?(?:#(?P<ID>[\w\-]+))?(?P<Classes>(?:\.[\w\-]+)*)(?P<Attributes>(?:\[[^\]]+\])*)$`, qArr[0])

// 	if len(a) > 0 {
// 		r := a[0]

// 		if s := r[1]; s != "" {
// 			newQ.tagName = s
// 		}
// 		if s := r[2]; s != "" {
// 			newQ.id = s
// 		}
// 		if s := r[3]; s != "" {
// 			a := strings.Split(s, ".")[1:]
// 			newQ.classList = a
// 		}
// 		if s := r[4]; s != "" {
// 			attrS := parseAttribute(s[1 : len(s)-1])
// 			newQ.attributes = attributes{}
// 			newQ.attributes[attrS.name] = attrS.value
// 		}
// 	}

// 	newQ.child = createQuery(qArr[1:]...)

// 	return &newQ
// }
