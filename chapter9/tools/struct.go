package tools

import (
	"errors"
	"fmt"
)

type c struct {
	Branch bool
}

func (c *c) example3() error {
	fmt.Println("in example3")
	if c.Branch {
		fmt.Println("branching code!")
		return errors.New("bad branch")
	}
	return nil
}
