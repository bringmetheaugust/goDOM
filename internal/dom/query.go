package dom

import (
	"regexp"
	"strings"

	"github.com/bringmetheaugust/goDOM/internal/errors"
)

type query struct {
	tagName   string
	id        string
	classList []string
	child     *query // next level query (ex in ".lol .lal" .lal is a next level query)
}

// Parse and get query as struct.
func parseQuery(qStr string) (*query, error) {
	if qStr == "" {
		return &query{}, errors.InvalidQuery{QueryStr: qStr}
	}

	queries := strings.Fields(qStr)

	return createQuery(queries...), nil
}

// Get query param from regExp
func matchedQueryParam(str string, q string) [][]string {
	reTag := regexp.MustCompile(str)

	return reTag.FindAllStringSubmatch(q, -1)
}

// Create query struct
func createQuery(qArr ...string) *query {
	if len(qArr) == 0 {
		return nil
	}

	q := qArr[0]
	newQ := query{}

	// get tagName
	matchedTag := matchedQueryParam(`^(\w+)(?:\.|#|$)`, q)
	if len(matchedTag) > 0 {
		newQ.tagName = matchedTag[0][1]
	}

	// get classes
	var classList []string
	matchedClasses := matchedQueryParam(`(?:\.)([\w-]+)`, q)
	for _, class := range matchedClasses {
		classList = append(classList, class[1])
	}
	newQ.classList = classList

	// get id
	matchedId := matchedQueryParam(`(?:#)([\w-]+)`, q)
	if len(matchedId) > 0 {
		newQ.id = matchedId[0][1]
	}
	newQ.child = createQuery(qArr[1:]...)

	return &newQ
}
