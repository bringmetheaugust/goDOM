package goDom

import "strings"

const (
	token_element = iota // HTML element
	token_content        // inner text
)

type token struct {
	_type int
	data  string
}

// Divide HTML to tokens. Each token is a HTML node.
// Uses as middle stream.
func tokenize(upStream chan string, downStream chan token) {
	var currentToken strings.Builder
	var isTrash bool

	for str := range upStream {
		for i := 0; i < len(str); i++ {
			switch str[i] {
			case '<':
				if currentToken.Len() > 0 && !isTrash { // new tag
					downStream <- token{token_element, currentToken.String()}
					currentToken.Reset()
				}

				currentToken.WriteByte('<')
			case '>':
				currentToken.WriteByte('>')

				switch t := currentToken.String(); {
				case strings.HasSuffix(t, "-->"), strings.HasPrefix(t, "<script"), strings.HasPrefix(t, "<style"): // comment is closing || script or style are opening
					downStream <- token{token_element, currentToken.String()}
					currentToken.Reset()

					isTrash = !isTrash
				case strings.HasSuffix(t, "</style>"): // style is closing
					content := currentToken.String()[:len(currentToken.String())-8]

					if content != "" {
						downStream <- token{token_content, content}
					}

					downStream <- token{token_element, "</style>"}
					currentToken.Reset()

					isTrash = false
				case strings.HasSuffix(t, "</script>"): // script is closing
					content := currentToken.String()[:len(currentToken.String())-9]

					if content != "" {
						downStream <- token{token_content, content}
					}

					downStream <- token{token_element, "</script>"}
					currentToken.Reset()

					isTrash = false
				case !isTrash: // simple tag closing
					downStream <- token{token_element, currentToken.String()}
					currentToken.Reset()
				}
			case '-':
				if currentToken.String() == "<!-" {
					isTrash = true
				}

				currentToken.WriteByte('-')
			default:
				// Concate lines with empty space if tag has multi lines in markup
				if i == 0 && currentToken.Len() > 0 {
					currentToken.WriteRune(' ')
				}

				currentToken.WriteByte(str[i])
			}
		}
	}

	close(downStream)
}
