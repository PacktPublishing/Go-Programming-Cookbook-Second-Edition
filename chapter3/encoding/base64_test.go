package encoding

import "testing"

func TestBase64Example(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Base64Example(); (err != nil) != tt.wantErr {
				t.Errorf("Base64Example() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBase64ExampleEncoder(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Base64ExampleEncoder(); (err != nil) != tt.wantErr {
				t.Errorf("Base64ExampleEncoder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
