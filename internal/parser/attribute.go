package parser

import (
	"strings"
)

type attribute struct {
	name  string
	value string
}

// Convert attribute from markup string to struct.
func parseAttribute(attr string) attribute {
	res := strings.Split(attr, "=")
	var attrVal string

	if len(res) > 1 {
		attrVal = strings.ReplaceAll(res[1], "'", "")
	}

	return attribute{
		name:  res[0],
		value: attrVal,
	}
}
