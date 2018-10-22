package filedirs

import "testing"

func TestOperate(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Operate(); (err != nil) != tt.wantErr {
				t.Errorf("Operate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
