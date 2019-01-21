package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter7/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
