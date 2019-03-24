package pool

import "testing"

func Test_hashWork(t *testing.T) {
	type args struct {
		wr WorkRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{WorkRequest{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashWork(tt.args.wr)
		})
	}
}

func Test_compareWork(t *testing.T) {
	type args struct {
		wr WorkRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{WorkRequest{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compareWork(tt.args.wr)
		})
	}
}
