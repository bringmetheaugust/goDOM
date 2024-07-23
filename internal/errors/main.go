package errors

import "fmt"

type NotFoundErr struct{}

func (e NotFoundErr) Error() string {
	return fmt.Sprintf("not found")
}
