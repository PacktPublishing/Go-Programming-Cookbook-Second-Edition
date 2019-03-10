package mocking

import "testing"

func TestRestorer_Restore(t *testing.T) {
	tests := []struct {
		name string
		r    Restorer
	}{
		{"base-case", Restorer(func() {})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Restore()
		})
	}
}

var a string

func TestPatch(t *testing.T) {
	type args struct {
		dest  interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{&a, "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Patch(tt.args.dest, tt.args.value)
		})
	}
}
