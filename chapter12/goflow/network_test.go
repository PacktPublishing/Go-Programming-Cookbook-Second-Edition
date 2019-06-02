package goflow

import "testing"

func TestNewEncodingApp(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := NewEncodingApp()
			if network == nil {
				t.Errorf("network unexpectedly nil")
			}
		})
	}
}
