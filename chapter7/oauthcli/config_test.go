package oauthcli

import (
	"context"
	"os"
	"reflect"
	"testing"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		name string
		want *oauth2.Config
	}{
		{"base-case",
			&oauth2.Config{
				ClientID:     os.Getenv("GITHUB_CLIENT"),
				ClientSecret: os.Getenv("GITHUB_SECRET"),
				Scopes:       []string{"repo", "user"},
				Endpoint:     github.Endpoint,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Setup(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Setup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetToken(t *testing.T) {
	type args struct {
		ctx  context.Context
		conf *oauth2.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{context.Background(), Setup()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetToken(tt.args.ctx, tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
