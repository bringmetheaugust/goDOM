package errors

import "fmt"

type InvalidRequest struct {
	Place string
}

func (e InvalidRequest) Error() string {
	return fmt.Sprintf("Invalid request or params to %v", e.Place)
}
