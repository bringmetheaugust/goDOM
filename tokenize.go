package goDom

import "strings"

// Divide HTML to tokens. Each token is a HTML node.
// Uses as middle stream.
func tokenize(upStream chan string, downStream chan string) {
	var currentToken strings.Builder
	var isTrash bool

	for str := range upStream {
		for i := 0; i < len(str); i++ {
			switch str[i] {
			case '<':
				if currentToken.Len() > 0 && !isTrash {
					downStream <- currentToken.String()
					currentToken.Reset()
				}

				currentToken.WriteByte('<')
			case '>':
				currentToken.WriteByte('>')

				switch t := currentToken.String(); {
				case strings.HasSuffix(t, "-->"), strings.HasPrefix(t, "<script"), strings.HasPrefix(t, "<style"): // comment is closing || script or style are opening
					downStream <- currentToken.String()
					currentToken.Reset()

					isTrash = !isTrash

					continue
				case strings.HasSuffix(t, "</style>"): // style is closing
					content := currentToken.String()[:len(currentToken.String())-8]

					if content != "" {
						downStream <- content
					}

					downStream <- "</style>"
					currentToken.Reset()

					isTrash = false

					continue
				case strings.HasSuffix(t, "</script>"): // script is closing
					content := currentToken.String()[:len(currentToken.String())-9]

					if content != "" {
						downStream <- content
					}

					downStream <- "</script>"
					currentToken.Reset()

					isTrash = false

					continue
				case !isTrash: // simple tag closing
					downStream <- currentToken.String()
					currentToken.Reset()

					continue
				}
			case '-':
				if currentToken.String() == "<!-" {
					isTrash = true
				}

				currentToken.WriteByte('-')
			default:
				currentToken.WriteByte(str[i])
			}
		}
	}

	close(downStream)
}
