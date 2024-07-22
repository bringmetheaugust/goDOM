package htmlparser

import "strings"

type tag struct {
	name       string
	attributes []attribute
}

// Parse HTML tag. Get tag name and rest attributes.
func parseTag(markup string) tag {
	var attributes []attribute

	tagStr := markup[1 : len(markup)-1]
	tagSplited := strings.Fields(tagStr)

	for _, attr := range tagSplited[1:] {
		attr := parseAttribute(attr)
		attributes = append(attributes, attr)
	}

	return tag{tagSplited[0], attributes}
}
