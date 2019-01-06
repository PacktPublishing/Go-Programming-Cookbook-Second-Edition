package main

import (
	"fmt"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter4/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()
}
