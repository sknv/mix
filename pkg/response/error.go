package response

import (
	"fmt"
)

const (
	MsgSmoke = "Что-то пошло не так"
)

// Error holds an error code, message and error itself.
type Error struct {
	Code     int
	Message  string
	Internal error
}

// NewError returns a response error.
func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// SetInternal sets an internal error to be logged.
func (e *Error) SetInternal(err error) *Error {
	e.Internal = err
	return e
}

func (e *Error) Error() string {
	return fmt.Sprintf("code = %d desc = %s", e.Code, e.Message)
}
