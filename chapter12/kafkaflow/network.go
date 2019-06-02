package kafkaflow

import "github.com/trustmaster/goflow"

// NewUpperApp wires together the compoents
func NewUpperApp() *goflow.Graph {
	u := goflow.NewGraph()

	u.Add("upper", new(Upper))
	u.Add("printer", new(Printer))

	u.Connect("upper", "Res", "printer", "Line")
	u.MapInPort("In", "upper", "Val")

	return u
}
