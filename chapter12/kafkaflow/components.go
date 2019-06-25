package kafkaflow

import (
	"fmt"
	"strings"

	flow "github.com/trustmaster/goflow"
)

// Upper upper cases the incoming
// stream
type Upper struct {
	Val <-chan string
	Res chan<- string
}

// Process loops over the input values and writes the upper
// case string version of them to Res
func (e *Upper) Process() {
	for val := range e.Val {
		e.Res <- strings.ToUpper(val)
	}
}

// Printer is a component for printing to stdout
type Printer struct {
	flow.Component
	Line <-chan string
}

// Process Prints the current line received
func (p *Printer) Process() {
	for line := range p.Line {
		fmt.Println(line)
	}
}
