package decorator

import "testing"

func TestSetup(t *testing.T) {

	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Setup()
		})
	}
}
