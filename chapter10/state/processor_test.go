package state

import (
	"context"
	"testing"
	"time"
)

func TestProcessor(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
	defer cancel()
	type args struct {
		ctx context.Context
		in  chan *WorkRequest
		out chan *WorkResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{ctx, make(chan *WorkRequest, 2), make(chan *WorkResponse, 2)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.in <- &WorkRequest{}
			go Processor(tt.args.ctx, tt.args.in, tt.args.out)
		})
	}
}
