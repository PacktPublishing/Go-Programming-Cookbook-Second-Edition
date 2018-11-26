package dataconv

import "testing"

func TestStrconv(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Strconv(); (err != nil) != tt.wantErr {
				t.Errorf("Strconv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
