package rest

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewAPIClient(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want *APIClient
	}{
		{"base-case", args{"username", "password"},
			&APIClient{
				Client: &http.Client{
					Transport: &APITransport{
						Transport: &http.Transport{},
						username:  "username",
						password:  "password",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAPIClient(tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPIClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIClient_GetGoogle(t *testing.T) {
	c := NewAPIClient("username", "password")
	tests := []struct {
		name    string
		c       *APIClient
		want    int
		wantErr bool
	}{
		{"base-case", c, 200, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetGoogle()
			if (err != nil) != tt.wantErr {
				t.Errorf("APIClient.GetGoogle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("APIClient.GetGoogle() = %v, want %v", got, tt.want)
			}
		})
	}
}
