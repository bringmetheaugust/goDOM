package errors

type InvalidQuery err

func (e InvalidQuery) Error() string {
	panic("unimplemented")
}
