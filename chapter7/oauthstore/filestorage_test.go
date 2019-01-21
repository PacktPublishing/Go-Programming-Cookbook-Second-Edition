package oauthstore

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"time"

	"golang.org/x/oauth2"
)

func TestFileStorage_GetToken(t *testing.T) {
	fname, _ := ioutil.TempFile(".", "example")
	defer os.Remove(fname.Name())
	tests := []struct {
		name    string
		f       *FileStorage
		want    *oauth2.Token
		wantErr bool
	}{
		{"base-case", &FileStorage{Path: fname.Name()}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.GetToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("FileStorage.GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileStorage.GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileStorage_SetToken(t *testing.T) {
	fname, _ := ioutil.TempFile(".", "example")
	defer os.Remove(fname.Name())
	type args struct {
		t *oauth2.Token
	}
	tests := []struct {
		name    string
		f       *FileStorage
		args    args
		wantErr bool
	}{
		{"base-case", &FileStorage{Path: fname.Name()}, args{&oauth2.Token{AccessToken: "123", Expiry: time.Now().Add(30 * time.Minute)}}, false},
		{"nil token", &FileStorage{Path: fname.Name()}, args{nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.SetToken(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("FileStorage.SetToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
