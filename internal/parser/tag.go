package parser

import (
	"goDOM/internal/dom"
	"regexp"
)

type tag struct {
	name       string
	attributes dom.Attributes
}

// Parse tokenized tag (without </>). Get tag name and rest attributes.
func parseTag(markup string) tag {
	attributes := make(dom.Attributes)

	re := regexp.MustCompile(`([^\s=]+='[^']*'|[^\s=]+)`)
	tagSplited := re.FindAllString(markup, -1)

	for _, attr := range tagSplited[1:] {
		attr := parseAttribute(attr)
		attributes[attr.name] = attr.value
	}

	return tag{tagSplited[0], attributes}
}
