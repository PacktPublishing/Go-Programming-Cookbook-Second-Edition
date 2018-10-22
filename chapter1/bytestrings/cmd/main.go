package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/bytestrings"

func main() {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	// each of these print to stdout
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
