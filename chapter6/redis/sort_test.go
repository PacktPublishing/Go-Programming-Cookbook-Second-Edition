package redis

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Sort(); (err != nil) != tt.wantErr {
				t.Errorf("Sort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
