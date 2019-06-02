package goflow

import (
	"encoding/base64"
	"fmt"
)

// Encoder base64 encodes all input
type Encoder struct {
	Val <-chan string
	Res chan<- string
}

// Process does the encoding then pushes the result onto Res
func (e *Encoder) Process() {
	for val := range e.Val {
		encoded := base64.StdEncoding.EncodeToString([]byte(val))
		e.Res <- fmt.Sprintf("%s => %s", val, encoded)
	}
}

// Printer is a component for printing to stdout
type Printer struct {
	Line <-chan string
}

// Process Prints the current line received
func (p *Printer) Process() {
	for line := range p.Line {
		fmt.Println(line)
	}
}
