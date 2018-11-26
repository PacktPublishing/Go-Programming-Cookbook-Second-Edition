package nulls

import "testing"

func TestPointerEncoding(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PointerEncoding(); (err != nil) != tt.wantErr {
				t.Errorf("PointerEncoding() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
