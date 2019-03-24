package pool

import (
	"context"
	"testing"
)

func TestDispatch(t *testing.T) {
	type args struct {
		numWorker int
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			canc, _, _ := Dispatch(tt.args.numWorker)
			defer canc()
		})
	}
}

func TestWorker(t *testing.T) {

	type args struct {
		id  int
		in  chan WorkRequest
		out chan WorkResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{1, make(chan WorkRequest, 2), make(chan WorkResponse)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()
			go Worker(ctx, tt.args.id, tt.args.in, tt.args.out)
			tt.args.in <- WorkRequest{}
			<-tt.args.out
		})
	}
}
