package htmlparser

import "strings"

type attribute struct {
	name  string
	value string
}

// Convert attribute from markup string to struct
func parseAttribute(attr string) attribute {
	res := strings.Split(attr, "=")

	return attribute{res[0], res[1]}
}
