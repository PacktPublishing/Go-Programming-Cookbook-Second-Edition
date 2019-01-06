package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// WrappedError demonstrates error wrapping and
// annotating an error
func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred in WrappedError")
}

// ErrorTyped is a error we can check against
type ErrorTyped struct {
	error
}

// Wrap shows what happens when we wrap an error
func Wrap() {
	e := errors.New("standard error")

	fmt.Println("Regular Error - ", WrappedError(e))

	fmt.Println("Typed Error - ", WrappedError(ErrorTyped{errors.New("typed error")}))

	fmt.Println("Nil -", WrappedError(nil))

}
