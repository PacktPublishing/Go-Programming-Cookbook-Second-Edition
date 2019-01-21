package oauthstore

import (
	"context"
	"testing"

	"golang.org/x/oauth2"
)

func TestGetToken(t *testing.T) {
	conf := Config{&oauth2.Config{}, &FileStorage{}}
	ctx := context.Background()
	type args struct {
		ctx  context.Context
		conf Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{ctx, conf}, true},
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
