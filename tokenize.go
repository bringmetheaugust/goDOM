package goDom

import "strings"

// Divide HTML to tokens. Each token is a HTML node.
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
			currentToken.WriteByte('>')

			switch t := currentToken.String(); {
			case strings.HasSuffix(t, "-->"), strings.HasPrefix(t, "<script"), strings.HasPrefix(t, "<style"): // comment is closing || script or style are opening
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()

				isTrash = !isTrash

				continue
			case strings.HasSuffix(t, "</style>"): // style is closing
				content := currentToken.String()[:len(currentToken.String())-8]

				if content != "" {
					tokens = append(tokens, content)
				}

				tokens = append(tokens, "</style>")
				currentToken.Reset()

				isTrash = false

				continue
			case strings.HasSuffix(t, "</script>"): // script is closing
				content := currentToken.String()[:len(currentToken.String())-9]

				if content != "" {
					tokens = append(tokens, content)
				}

				tokens = append(tokens, "</script>")
				currentToken.Reset()

				isTrash = false

				continue
			case !isTrash: // simple tag closing
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()

				continue
			}
		case '-':
			if currentToken.String() == "<!-" {
				isTrash = true
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
