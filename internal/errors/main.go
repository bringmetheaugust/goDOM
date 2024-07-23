package errors

import "fmt"

type NotFound struct {
	Query string
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("%v not found", e.Query)
}
