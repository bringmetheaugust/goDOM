package parser

import (
	"goDOM/internal/dom"
	"strings"
)

type tag struct {
	name       string
	attributes []dom.Attribute
}

// Parse HTML tag. Get tag name and rest attributes.
func parseTag(markup string) tag {
	var attributes []dom.Attribute

	tagStr := markup[1 : len(markup)-1]
	tagSplited := strings.Fields(tagStr)

	for _, attr := range tagSplited[1:] {
		attr := parseAttribute(attr)
		attributes = append(attributes, attr)
	}

	return tag{tagSplited[0], attributes}
}
