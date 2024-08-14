package goDom

import "fmt"

type notFoundErr struct {
	Params string
	Msg    string
}

//	type invalidRequestErr struct {
//		Place string
//	}
type invalidQueryErr struct {
	QueryStr string
	Msg      string
}

func (e notFoundErr) Error() string {
	return fmt.Sprintf("Not found with params: %v. %v ", e.Params, e.Msg)
}

// func (e invalidRequestErr) Error() string {
// 	return fmt.Sprintf("Invalid request or params to %v", e.Place)
// }

func (e invalidQueryErr) Error() string {
	return fmt.Sprintf("invalid query: %v. %v ", e.QueryStr, e.Msg)
}
