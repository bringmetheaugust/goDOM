package dom

import (
	"goDOM/internal/errors"
	"regexp"
	"strings"
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
		return &query{}, errors.InvalidQuery{}
	}

	queries := strings.Fields(qStr)

	return createQuery(queries...), nil
}

// TODO refactoring
// create query struct
func createQuery(qArr ...string) *query {
	if len(qArr) == 0 {
		return nil
	}

	q := qArr[0]
	newQ := query{}

	// get tagName
	reTag := regexp.MustCompile(`^(\w+)(?:\.|#|$)`)
	matchedTag := reTag.FindAllStringSubmatch(q, -1)

	if len(matchedTag) > 0 {
		newQ.tagName = matchedTag[0][1]
	}

	// get classes
	var classList []string
	reCLass := regexp.MustCompile(`(?:\.)([\w-]+)`)
	matchedClasses := reCLass.FindAllStringSubmatch(q, -1)

	for _, class := range matchedClasses {
		classList = append(classList, class[1])
	}

	newQ.classList = classList

	// get id
	reId := regexp.MustCompile(`(?:#)([\w-]+)`)
	matchedId := reId.FindAllStringSubmatch(q, -1)

	if len(matchedId) > 0 {
		newQ.id = matchedId[0][1]
	}

	newQ.child = createQuery(qArr[1:]...)

	return &newQ
}
