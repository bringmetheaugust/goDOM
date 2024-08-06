package goDom

import "fmt"

type notFoundErr struct{}
type invalidRequestErr struct {
	Place string
}
type invalidQueryErr struct {
	QueryStr string
}

func (e notFoundErr) Error() string {
	return "Not found"
}

func (e invalidRequestErr) Error() string {
	return fmt.Sprintf("Invalid request or params to %v", e.Place)
}

func (e invalidQueryErr) Error() string {
	return fmt.Sprintf("invalid query: %v", e.QueryStr)
}
