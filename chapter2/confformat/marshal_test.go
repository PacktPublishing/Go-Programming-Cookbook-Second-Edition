package confformat

import "testing"

func TestMarshalAll(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MarshalAll(); (err != nil) != tt.wantErr {
				t.Errorf("MarshalAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
