package pool

import "testing"

func TestProcess(t *testing.T) {
	type args struct {
		wr WorkRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{"Hash", args{WorkRequest{Op: Hash}}},
		{"Compare", args{WorkRequest{Op: Compare}}},
		{"Other", args{WorkRequest{Op: op("test")}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Process(tt.args.wr)
		})
	}
}
