package context

import (
	"context"
	"reflect"
	"testing"
)

func TestSetup(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, timeoutKey, "timeout exceeded")
	ctx = context.WithValue(ctx, deadlineKey, "deadline exceeded")

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{"base-case", args{context.Background()}, ctx},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Setup(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Setup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValue(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, timeoutKey, "timeout exceeded")

	type args struct {
		ctx context.Context
		k   key
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"has key", args{ctx, timeoutKey}, "timeout exceeded"},
		{"no key", args{ctx, deadlineKey}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValue(tt.args.ctx, tt.args.k); got != tt.want {
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
