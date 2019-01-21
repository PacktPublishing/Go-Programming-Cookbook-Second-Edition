package oauthcli

import (
	"net/http"
	"testing"
)

func TestGetUsers(t *testing.T) {
	type args struct {
		client *http.Client
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{&http.Client{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetUsers(tt.args.client); (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
