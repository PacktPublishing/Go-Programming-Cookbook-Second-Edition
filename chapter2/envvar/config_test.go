package envvar

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	c := struct {
		Example string `json:"example"`
	}{}
	type args struct {
		path      string
		envPrefix string
		config    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{"123", "", &c}, true},
		{"no file", args{"", "", &c}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadConfig(tt.args.path, tt.args.envPrefix, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoadFile(t *testing.T) {
	c := struct {
		Example string `json:"example"`
	}{}

	tf, err := ioutil.TempFile("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())

	if _, err := tf.Write(bytes.NewBufferString("{}").Bytes()); err != nil {
		t.Errorf("Error writing into temp file, err: %v ", err)
	}

	type args struct {
		path   string
		config interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{"", &c}, true},
		{"base-case", args{tf.Name(), &c}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadFile(tt.args.path, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("LoadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
