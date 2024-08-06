package goDom

import (
	"regexp"
	"strings"
)

type tag struct {
	name       string
	attributes attributes
}

// Parse tokenized tag. Get tag name and rest attributes.
func parseTag(markup string) tag {
	var tagStr string
	attributes := make(attributes)

	if strings.HasSuffix(markup, "/>") {
		tagStr = markup[1 : len(markup)-2]
	} else {
		tagStr = markup[1 : len(markup)-1]
	}

	re := regexp.MustCompile(`([^\s=]+='[^']*'|[^\s=]+)`)
	tagSplited := re.FindAllString(tagStr, -1)

	for _, attr := range tagSplited[1:] {
		attr := parseAttribute(attr)
		attributes[attr.name] = attr.value
	}

	if len(attributes) == 0 {
		attributes = nil
	}

	return tag{tagSplited[0], attributes}
}
