package parser

import (
	"goDOM/internal/vdom"
	"strings"
)

// Convert attribute from markup string to struct
func parseAttribute(attr string) vdom.Attribute {
	res := strings.Split(attr, "=")

	return vdom.Attribute{res[0], res[1]}
}
