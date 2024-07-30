package parser

import "strings"

// Divide HTML to tokens. Each token is a HTML node.
func tokenize(input string) []string {
	var tokens []string
	var currentToken strings.Builder

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '<':
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}

			currentToken.WriteByte('<')
		case '>':
			currentToken.WriteByte('>')
			tokens = append(tokens, currentToken.String())
			currentToken.Reset()
		default:
			currentToken.WriteByte(input[i])
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}
