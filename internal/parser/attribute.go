package parser

import (
	"goDOM/internal/dom"
	"strings"
)

// Convert attribute from markup string to struct
func parseAttribute(attr string) dom.Attribute {
	res := strings.Split(attr, "=")

	return dom.Attribute{
		Name:  res[0],
		Value: strings.ReplaceAll(res[1], "'", ""),
	}
}
