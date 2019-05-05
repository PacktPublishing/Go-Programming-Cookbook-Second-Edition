package consensus

import "testing"

func TestConfig(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Config(tt.args.num)
		})
	}
}
