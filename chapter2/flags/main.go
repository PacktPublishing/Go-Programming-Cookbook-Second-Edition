package main

import (
	"flag"
	"fmt"
)

func main() {
	// initialize our setup
	c := Config{}
	c.Setup()

	// generally call this from main
	flag.Parse()

	fmt.Println(c.GetMessage())
}
