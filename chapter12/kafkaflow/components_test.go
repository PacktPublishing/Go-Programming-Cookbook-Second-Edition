package kafkaflow

import (
	"testing"

	flow "github.com/trustmaster/goflow"
)

func TestUpper_OnVal(t *testing.T) {
	type fields struct {
		Component flow.Component
		Val       <-chan string
		Res       chan<- string
	}
	type args struct {
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base-case", fields{flow.Component{}, make(chan string, 2), make(chan string, 2)}, args{"test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Upper{
				Component: tt.fields.Component,
				Val:       tt.fields.Val,
				Res:       tt.fields.Res,
			}
			e.OnVal(tt.args.val)
		})
	}
}

func TestPrinter_OnLine(t *testing.T) {
	type fields struct {
		Component flow.Component
		Line      <-chan string
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base-case", fields{flow.Component{}, make(chan string, 2)}, args{"test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Printer{
				Component: tt.fields.Component,
				Line:      tt.fields.Line,
			}
			p.OnLine(tt.args.line)
		})
	}
}
