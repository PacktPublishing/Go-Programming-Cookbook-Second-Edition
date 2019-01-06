package errwrap

import (
	"errors"
	"testing"
)

func TestWrappedError(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"nil", args{nil}, false},
		{"error", args{errors.New("real")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WrappedError(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("WrappedError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Wrap()
		})
	}
}
