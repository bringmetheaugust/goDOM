package parser

import (
	"strings"
)

// Divide HTML to tokens. Each token is a HTML node.
func tokenize(input string) []string {
	var tokens []string
	var currentToken strings.Builder
	var isComment bool

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '<':
			if currentToken.Len() > 0 && !isComment {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}

			currentToken.WriteByte('<')
		case '>': // check if tag is closing or it's a part of comment content
			if !isComment || (isComment && strings.HasSuffix(currentToken.String(), "--")) { // tag or comment are closing
				currentToken.WriteByte('>')
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()

				if isComment {
					isComment = false
				}

				continue
			}

			currentToken.WriteByte('>')
		case '-': // for comments (ex "<!-- <div></div> -->")
			if strings.Compare(currentToken.String(), "<!-") == 0 {
				isComment = true
			}

			currentToken.WriteByte('-')
		default:
			currentToken.WriteByte(input[i])
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}
