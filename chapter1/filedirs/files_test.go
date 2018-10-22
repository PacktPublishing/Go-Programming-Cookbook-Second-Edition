package filedirs

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCapitalizer(t *testing.T) {
	f1, err := ioutil.TempFile(".", "tmp")
	if err != nil {
		t.Error(err)
	}
	defer f1.Close()
	defer os.Remove(f1.Name())

	f2, err := ioutil.TempFile(".", "tmp")
	if err != nil {
		t.Error(err)
	}
	defer f2.Close()
	defer os.Remove(f2.Name())

	type args struct {
		f1 *os.File
		f2 *os.File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "base-case",
			args:    args{f1, f2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Capitalizer(tt.args.f1, tt.args.f2); (err != nil) != tt.wantErr {
				t.Errorf("Capitalizer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCapitalizerExample(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CapitalizerExample(); (err != nil) != tt.wantErr {
				t.Errorf("CapitalizerExample() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
