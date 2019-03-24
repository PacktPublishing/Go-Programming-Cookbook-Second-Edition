package channels

import (
	"context"
	"testing"
	"time"
)

func TestPrinter(t *testing.T) {

	type args struct {
		ctx context.Context
		ch  chan string
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{context.Background(), make(chan string)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(tt.args.ctx)
			defer cancel()
			go Printer(ctx, tt.args.ch)
			time.Sleep(200 * time.Millisecond)
			tt.args.ch <- "test"
		})
	}
}
