package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter7/decorator"

func main() {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}
