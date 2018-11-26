package encoding

import "testing"

func TestGobExample(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GobExample(); (err != nil) != tt.wantErr {
				t.Errorf("GobExample() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
