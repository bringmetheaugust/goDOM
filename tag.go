package goDom

import (
	"regexp"
	"strings"
)

type tag struct {
	name        string
	attributes  attributes
	selfClosing bool
}

// Parse tokenized tag. Get tag name and rest attributes.
func parseTag(markup string) tag {
	var tagStr string
	var newTag tag
	attributes := make(attributes)

	if strings.HasSuffix(markup, "/>") {
		newTag.selfClosing = true
		tagStr = markup[1 : len(markup)-2]
	} else {
		tagStr = markup[1 : len(markup)-1]
	}

	re := regexp.MustCompile(`([^\s=]+='[^']*'|[^\s=]+)`)
	tagSplited := re.FindAllString(tagStr, -1)
	newTag.name = tagSplited[0]

	for _, attr := range tagSplited[1:] {
		attr := parseAttribute(attr)
		attributes[attr.name] = attr.value
	}

	if len(attributes) != 0 {
		newTag.attributes = attributes
	}

	return newTag
}
