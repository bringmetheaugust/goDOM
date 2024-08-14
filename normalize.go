package goDom

import (
	"bufio"
	"strings"
)

// Prepare HTML for parsing.
// Remove endlines (\n), tabs (\t), empty lines.
// Replace " to '.
// Uses as upstream.
func normalize(markup string, ch chan string) {
	scanner := bufio.NewScanner(strings.NewReader(markup))

	for scanner.Scan() {
		str := scanner.Text()
		fmtStr_0 := strings.TrimSpace(str)

		if fmtStr_0 == "" {
			continue
		}

		fmtStr_1 := strings.ReplaceAll(fmtStr_0, `"`, `'`)

		ch <- fmtStr_1
	}

	close(ch)
}
