package errors

import "fmt"

type InvalidQuery struct {
	QueryStr string
}

func (e InvalidQuery) Error() string {
	return fmt.Sprintf("invalid query: %v", e.QueryStr)
}
