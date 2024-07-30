package parser

import "strings"

// Prepare HTML for parsing. Remove/replace some chars.
func normalize(markup string) string {
	markup = strings.ReplaceAll(markup, `"`, `'`) // cause string() shielding /"
	markup = strings.TrimSpace(markup)

	return markup
}
