package parser

import (
	"strings"
)

// Divide HTML to tokens. Each token is a HTML node.
// Ignore <!Tag>, comments and script's tags.
func tokenize(input string) []string {
	var tokens []string
	var currentToken strings.Builder
	var isTrash bool

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '<':
			if currentToken.Len() > 0 && !isTrash {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}

			currentToken.WriteByte('<')
		case '>':
			cT := currentToken.String()

			// close <! || close comment || script ends || style ends
			if (isTrash && strings.HasPrefix(cT, "<!") && !strings.HasPrefix(cT, "<!--")) || (isTrash && strings.HasSuffix(cT, "--")) || strings.HasSuffix(cT, "</script") || strings.HasSuffix(cT, "</style") {
				currentToken.Reset()
				isTrash = false

				continue
			}

			// open script or style
			if cT == "<script" || strings.HasPrefix(cT, "<style") {
				isTrash = true

				continue
			}

			// simple tag closing || opening
			if !isTrash {
				currentToken.WriteByte('>')
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()

				continue
			}

			currentToken.WriteByte('>')
		case '!':
			if currentToken.String() == "<" {
				isTrash = true
			}

			currentToken.WriteByte('!')
		default:
			currentToken.WriteByte(input[i])
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}
