package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter3/dataconv"

func main() {
	dataconv.ShowConv()
	if err := dataconv.Strconv(); err != nil {
		panic(err)
	}
	dataconv.Interfaces()
}
