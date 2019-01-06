package main

import (
	"fmt"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter4/basicerrors"
)

func main() {
	basicerrors.BasicErrors()

	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
