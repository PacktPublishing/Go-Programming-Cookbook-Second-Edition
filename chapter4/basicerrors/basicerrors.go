package basicerrors

import (
	"errors"
	"fmt"
)

// ErrorValue is a way to make a package level
// error to check against. I.e. if err == ErrorValue
var ErrorValue = errors.New("this is a typed error")

// TypedError is a way to make an error type
// you can do err.(type) == ErrorValue
type TypedError struct {
	error
}

//BasicErrors demonstrates some ways to create errors
func BasicErrors() {
	err := errors.New("this is a quick and easy way to create an error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("an error occurred: %s", "something")
	fmt.Println("fmt.Errorf: ", err)

	err = ErrorValue
	fmt.Println("value error: ", err)

	err = TypedError{errors.New("typed error")}
	fmt.Println("typed error: ", err)

}
