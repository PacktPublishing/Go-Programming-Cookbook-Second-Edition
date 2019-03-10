package tools

import (
	"fmt"
)

func example() error {
	fmt.Println("in example")
	return nil
}

var example2 = func() int {
	fmt.Println("in example2")
	return 10
}
