package goflow

import (
	"github.com/trustmaster/goflow"
)

// NewEncodingApp wires together the components
func NewEncodingApp() *goflow.Graph {
	e := goflow.NewGraph()

	// define component types
	e.Add("encoder", new(Encoder))
	e.Add("printer", new(Printer))

	// connect the components using channels
	e.Connect("encoder", "Res", "printer", "Line")

	// map the in channel to Val, which is
	// tied to OnVal function
	e.MapInPort("In", "encoder", "Val")

	return e
}
