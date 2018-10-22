package interfaces

import "testing"

func TestPipeExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PipeExample()
		})
	}
}
