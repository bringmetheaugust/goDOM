package parser

import (
	"goDOM/internal/dom"
	"strings"
)

// Convert attribute from markup string to struct.
func parseAttribute(attr string) dom.Attribute {
	res := strings.Split(attr, "=")
	var attrVal string

	if len(res) > 1 {
		attrVal = strings.ReplaceAll(res[1], "'", "")
	}

	return dom.Attribute{
		Name:  res[0],
		Value: attrVal,
	}
}
