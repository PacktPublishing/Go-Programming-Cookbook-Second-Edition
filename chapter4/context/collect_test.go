package context

import (
	"context"
	"testing"
)

func TestInitialize(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Initialize()
		})
	}
}

func Test_gatherName(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{context.Background()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gatherName(tt.args.ctx)
		})
	}
}

func Test_gatherLocation(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{context.Background()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gatherLocation(tt.args.ctx)
		})
	}
}
