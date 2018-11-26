package math

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	b := big.NewInt(1)
	b.SetString("573147844013817084101", 10)
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"n = 0", args{0}, big.NewInt(1)},
		{"n = 1", args{1}, big.NewInt(1)},
		{"n = 2", args{2}, big.NewInt(2)},
		{"n = 3", args{3}, big.NewInt(3)},
		{"n = 100", args{100}, b},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
