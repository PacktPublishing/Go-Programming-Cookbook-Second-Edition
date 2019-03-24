package state

import (
	"errors"
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {
	type args struct {
		wr *WorkRequest
	}
	tests := []struct {
		name string
		args args
		want *WorkResponse
	}{
		{"add", args{&WorkRequest{Add, 1, 2}}, &WorkResponse{Result: 3}},
		{"subtract", args{&WorkRequest{Subtract, 1, 2}}, &WorkResponse{Result: -1}},
		{"multiply", args{&WorkRequest{Multiply, 3, 3}}, &WorkResponse{Result: 9}},
		{"divide", args{&WorkRequest{Divide, 3, 3}}, &WorkResponse{Result: 1}},
		{"divide by zero", args{&WorkRequest{Divide, 3, 0}}, &WorkResponse{Err: errors.New("divide by 0")}},
		{"bad op", args{&WorkRequest{op("test"), 3, 0}}, &WorkResponse{Err: errors.New("unsupported operation")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want.Wr = tt.args.wr
			if got := Process(tt.args.wr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() = %v, want %v", got, tt.want)
			}
		})
	}
}
