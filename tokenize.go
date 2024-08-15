package goDom

import "strings"

const (
	node_element = iota // HTML element
	node_text           // inner text
	node_meta           // doctype
	node_comment        // comment
)

type token struct {
	_type     int
	data      string
	tag       tag
	isClosing bool
}

// Divide HTML to tokens. Each token is a HTML node. Uses as middle stream.
func tokenize(upStream chan string, downStream chan token) {
	var b strings.Builder
	var isTrash bool // uses for comment/script/style to disable tokenization of their content

	for str := range upStream {
	strLoop:
		for i := 0; i < len(str); i++ {
			switch str[i] {
			case '<': // new tag
				if b.Len() > 0 && !isTrash { // text node
					downStream <- token{_type: node_text, data: b.String()}
					b.Reset()
				}

				b.WriteByte('<')
			case '>':
				b.WriteByte('>')

				switch t := b.String(); {
				case strings.HasSuffix(t, "-->"): // comment is closing
					downStream <- token{_type: node_comment, data: b.String()}
					isTrash = false
				case strings.HasPrefix(t, "<script"), strings.HasPrefix(t, "<style"): // script or style are opening
					tag := parseTag(t)
					downStream <- token{_type: node_element, tag: tag}
					isTrash = true
				case strings.HasSuffix(t, "</style>"): // style is closing
					// this and the next case (with "</script>") have allmost the same logic
					b.WriteRune('>')
					fallthrough
				case strings.HasSuffix(t, "</script>"): // script is closing
					content := b.String()[:len(b.String())-9]

					if content != "" {
						downStream <- token{_type: node_text, data: content}
					}

					downStream <- token{_type: node_element, isClosing: true}
					isTrash = false
				case strings.HasPrefix(t, "<!") && !isTrash: // doctype ends
					downStream <- token{_type: node_meta, data: b.String()}
				case strings.HasPrefix(t, "</") && !isTrash: // element is slosing
					downStream <- token{_type: node_element, isClosing: true}
				case !isTrash: // tag is closing
					tag := parseTag(b.String())
					downStream <- token{_type: node_element, tag: tag}
				default: // trash
					continue strLoop
				}

				b.Reset()
			case '-':
				if b.String() == "<!-" { // if comment is starting
					isTrash = true
				}

				b.WriteByte('-')
			default:
				// Concate lines with empty space if tag has multi lines in markup
				if i == 0 && b.Len() > 0 {
					b.WriteRune(' ')
				}

				b.WriteByte(str[i])
			}
		}
	}

	close(downStream)
}
