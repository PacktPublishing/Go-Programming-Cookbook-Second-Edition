package templates

import "testing"

func TestHTMLDifferences(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := HTMLDifferences(); (err != nil) != tt.wantErr {
				t.Errorf("HTMLDifferences() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
