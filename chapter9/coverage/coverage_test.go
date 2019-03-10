package main

import "testing"

func TestCoverage(t *testing.T) {
	type args struct {
		condition bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"no condition", args{true}, true},
		{"condition", args{false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Coverage(tt.args.condition); (err != nil) != tt.wantErr {
				t.Errorf("Coverage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
