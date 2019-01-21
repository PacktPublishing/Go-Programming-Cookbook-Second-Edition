package client

import (
	"net/http"
	"testing"
)

func TestDoOps(t *testing.T) {
	c := Setup(true, true)
	type args struct {
		c *http.Client
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{c}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DoOps(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DoOps() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefaultGetGolang(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DefaultGetGolang(); (err != nil) != tt.wantErr {
				t.Errorf("DefaultGetGolang() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
