package goDom

import (
	"fmt"
	"regexp"
	"strings"
)

type query struct {
	tagName string
	id      string
	classList
	attributes
	operator selectorOperator // if query is selector operator. "" if not
	child    *query           // next level query (ex in ".lol .lal" .lal is a next level query)
}

type queries []query

type selectorOperator string

const (
	query_operator_all    selectorOperator = "*"
	query_operator_parent selectorOperator = ">"
	query_operator_after  selectorOperator = "+"
	query_operator_hz     selectorOperator = "~"
)

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
	var matchedTag, matchedClasses, matchedId, matchedAttrs [][]string
	var classList classList

	// if selector operator
	switch o := selectorOperator(q); o {
	case query_operator_all:
		newQ.operator = selectorOperator(o)
		goto nextQuery
	case query_operator_parent, query_operator_after, query_operator_hz:
		panic(fmt.Sprintf("%v operator is not jet supported.", o))
	}

	// * get tagName
	matchedTag = matchedQueryParam(`^(\w+)(?:\.|\[|#|$)`, q)
	if len(matchedTag) > 0 {
		newQ.tagName = matchedTag[0][1]
	}

	// * get classes
	matchedClasses = matchedQueryParam(`(?:\.)([\w-]+)`, q)
	for _, class := range matchedClasses {
		classList = append(classList, class[1])
	}
	newQ.classList = classList

	// * get id
	matchedId = matchedQueryParam(`(?:#)([\w-]+)`, q)
	if len(matchedId) > 0 {
		newQ.id = matchedId[0][1]
	}

	// * get attributes
	matchedAttrs = matchedQueryParam(`\[(\w+=[\'"][^\'"]*[\'"]|[\w-]+)\]`, q)
	if len(matchedAttrs) > 0 {
		newQ.attributes = attributes{}

		for _, m := range matchedAttrs {
			res := strings.Split(m[1], "=")
			var attrVal string

			if len(res) > 1 {
				attrVal = strings.ReplaceAll(res[1], "'", "")
			}

			newQ.attributes[res[0]] = attrVal
		}
	}

nextQuery:
	newQ.child = createQuery(qArr[1:]...)

	return &newQ
}

// Get query param from regExp
func matchedQueryParam(str string, q string) [][]string {
	reTag := regexp.MustCompile(str)

	return reTag.FindAllStringSubmatch(q, -1)
}
