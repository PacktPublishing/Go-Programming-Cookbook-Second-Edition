package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	c := MenuConf{}
	menu := c.SetupMenu()

	if err := menu.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing params %s, error: %v", os.Args[1:], err)
		return
	}

	// we use arguments to switch between commands
	// flags are also an argument
	if len(os.Args) > 1 {
		// we don't care about case
		switch strings.ToLower(os.Args[1]) {
		case "version":
			c.Version()
		case "greet":
			f := c.GetSubMenu()
			if len(os.Args) < 3 {
				f.Usage()
				return
			}
			if len(os.Args) > 3 {
				if err := f.Parse(os.Args[3:]); err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing params %s, error: %v", os.Args[3:], err)
					return
				}

			}
			c.Greet(os.Args[2])

		default:
			fmt.Println("Invalid command")
			menu.Usage()
			return
		}
	} else {
		menu.Usage()
		return
	}
}
