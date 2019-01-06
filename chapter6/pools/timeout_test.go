package pools

import "testing"

func TestExecWithTimeout(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExecWithTimeout(); (err != nil) != tt.wantErr {
				t.Errorf("ExecWithTimeout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
