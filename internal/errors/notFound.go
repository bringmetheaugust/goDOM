package errors

import "fmt"

type NotFound err

func (e NotFound) Error() string {
	return fmt.Sprintf("%v not found", e.query)
}
