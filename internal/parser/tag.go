package parser

import (
	"regexp"

	"github.com/bringmetheaugust/goDOM/internal/dom"
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

	if len(attributes) == 0 {
		attributes = nil
	}

	return tag{tagSplited[0], attributes}
}
