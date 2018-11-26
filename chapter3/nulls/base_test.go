package nulls

import "testing"

func TestBaseEncoding(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BaseEncoding(); (err != nil) != tt.wantErr {
				t.Errorf("BaseEncoding() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
