package messages

import "fmt"

type ErrorMessage string

func (e ErrorMessage) AsError(str ...string) error {
	res := string(e)
	for _, s := range str {
		res = fmt.Sprintf("%s %s", s, res)
	}
	return fmt.Errorf(res)
}

const (
	ALREADY_EXISTS  ErrorMessage = "already exists"
	DOES_NOT_EXISTS ErrorMessage = "does not exists"
	NOT_FOUND       ErrorMessage = "not found"
	IS_INVALID      ErrorMessage = "is invalid"
)
