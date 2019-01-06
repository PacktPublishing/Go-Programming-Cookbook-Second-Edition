package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// Unwrap will unwrap an error and do
// type assertion to it
func Unwrap() {

	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Println("wrapped error: ", err)

	// we can handle many error types
	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("a typed error occurred: ", err)
	default:
		fmt.Println("an unknown error occurred")
	}
}

// StackTrace will print all the stack for
// the error
func StackTrace() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Printf("%+v\n", err)
}
