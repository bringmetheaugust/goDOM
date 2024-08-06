package goDom

import (
	"bufio"
	"strings"
)

// Prepare HTML for parsing.
// Remove endlines (\n), tabs (\t), empty lines.
// Replace " to '.
func normalize(markup string) string {
	scanner := bufio.NewScanner(strings.NewReader(markup))
	var b strings.Builder

	for scanner.Scan() {
		str := scanner.Text()
		b.WriteString(strings.TrimSpace(str))
	}
	markup = b.String()
	markup = strings.ReplaceAll(markup, `"`, `'`) // cause string() shielding /"

	return markup
}
