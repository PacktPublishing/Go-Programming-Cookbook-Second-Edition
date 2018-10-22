package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter1/filedirs"

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	if err := filedirs.CapitalizerExample(); err != nil {
		panic(err)
	}
}
