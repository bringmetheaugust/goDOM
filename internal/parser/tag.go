package parser

import (
	"goDOM/internal/dom"
	"regexp"
)

type tag struct {
	name       string
	attributes []dom.Attribute
}

// Parse HTML tag. Get tag name and rest attributes.
func parseTag(markup string) tag {
	var attributes []dom.Attribute

	tagStr := markup[1 : len(markup)-1]
	re := regexp.MustCompile(`([^\s=]+='[^']*'|[^\s=]+)`)
	tagSplited := re.FindAllString(tagStr, -1)

	for _, attr := range tagSplited[1:] {
		attr := parseAttribute(attr)
		attributes = append(attributes, attr)
	}

	return tag{tagSplited[0], attributes}
}
