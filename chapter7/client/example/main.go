package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter7/client"

func main() {
	// secure and op!
	cli := client.Setup(true, false)

	if err := client.DefaultGetGolang(); err != nil {
		panic(err)
	}

	if err := client.DoOps(cli); err != nil {
		panic(err)
	}

	c := client.Controller{Client: cli}
	if err := c.DoOps(); err != nil {
		panic(err)
	}

	// secure and noop
	// also modifies default
	client.Setup(true, true)

	if err := client.DefaultGetGolang(); err != nil {
		panic(err)
	}
}
