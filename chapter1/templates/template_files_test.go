package templates

import "testing"
import "os"

func TestCreateTemplate(t *testing.T) {
	type args struct {
		path string
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{"test-will-delete", "hello"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTemplate(tt.args.path, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("CreateTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	os.RemoveAll("test-will-delete")
}

func TestInitTemplates(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitTemplates(); (err != nil) != tt.wantErr {
				t.Errorf("InitTemplates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
