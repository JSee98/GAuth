package error

import (
	"fmt"
	"strings"
)

/*
	Custom error object. 
	Shall be bubbled up and if additional details need to be passed should be added into the details field.
	The error will panic at the root of the trace not from the bottom.
*/
type Error struct {
	message string
	details []string
	shouldPanic bool
}

func NewError(msg string) *Error {
	return &Error{
		message: msg,
		details: make([]string, 0),
		shouldPanic: false,
	}
}

func NewErrorEmpty() *Error {
	return &Error{
		message: "",
		details: make([]string, 0),
		shouldPanic: false,
	}
}

func (e *Error) ParseToError() error {
	baseError := e.message

	hasDetails := len(e.details)

	if hasDetails>0{
		csl := strings.Join(e.details, ",")
		return fmt.Errorf("%s\nDetailed list of errors: %s\n", baseError, csl)
	}
	return fmt.Errorf("%s", baseError)
}

func (e *Error) AddDetails(detail string) {
	e.details = append(e.details,detail)
}

func (e *Error) SetPanic() {
	e.shouldPanic=true
}

func (e *Error) SetMessage(msg string) {
	e.message = msg
}

// returns an error or nil based on if the error was initialized at all
func (e *Error) Error() *Error {
	if len(e.details)>0 || e.message !="" || e.shouldPanic {
		return e
	}
	return nil
}