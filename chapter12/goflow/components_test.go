package goflow

import (
	"testing"
)

func TestEncoder_Process(t *testing.T) {
	type fields struct {
		Val chan string
		Res chan string
	}
	type args struct {
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result string
	}{
		{"base-case", fields{Val: make(chan string, 2), Res: make(chan string, 2)}, args{"test"}, "test => dGVzdA=="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				Val: tt.fields.Val,
				Res: tt.fields.Res,
			}
			go e.Process()
			tt.fields.Val <- "test"
			if got, want := <-tt.fields.Res, tt.result; got != want {
				t.Errorf("got %s; want %s", got, want)
			}
		})
	}
}

func TestPrinter_Process(t *testing.T) {
	type fields struct {
		Line chan string
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base-case", fields{make(chan string, 2)}, args{"test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Printer{
				Line: tt.fields.Line,
			}
			tt.fields.Line <- "test"
			go p.Process()
		})
	}
}
