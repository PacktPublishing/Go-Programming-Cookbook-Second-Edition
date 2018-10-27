package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter1/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
