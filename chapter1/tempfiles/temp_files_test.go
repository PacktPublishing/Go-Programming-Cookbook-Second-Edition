package tempfiles

import "testing"

func TestWorkWithTemp(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WorkWithTemp(); (err != nil) != tt.wantErr {
				t.Errorf("WorkWithTemp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
