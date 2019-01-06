package basicerrors

import (
	"fmt"
)

// CustomError is a struct that will implement
// the Error() interface
type CustomError struct {
	Result string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("there was an error; %s was the result", c.Result)
}

// SomeFunc returns an error
func SomeFunc() error {
	c := CustomError{Result: "this"}
	return c
}
