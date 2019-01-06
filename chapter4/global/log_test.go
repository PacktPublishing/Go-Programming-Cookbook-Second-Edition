package global

import "testing"

func TestUseLog(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UseLog(); (err != nil) != tt.wantErr {
				t.Errorf("UseLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
