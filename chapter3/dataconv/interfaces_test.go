package dataconv

import "testing"

func TestCheckType(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"string", args{"test"}},
		{"int", args{1}},
		{"bool", args{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckType(tt.args.s)
		})
	}
}

func TestInterfaces(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Interfaces()
		})
	}
}
