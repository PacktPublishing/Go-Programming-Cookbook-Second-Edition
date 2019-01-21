package oauthstore

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"golang.org/x/oauth2"
)

func TestConfig_Exchange(t *testing.T) {
	type args struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name    string
		c       *Config
		args    args
		wantErr bool
	}{
		{"base-case", &Config{&oauth2.Config{}, &FileStorage{}}, args{context.Background(), "123"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.c.Exchange(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Exchange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestConfig_TokenSource(t *testing.T) {
	conf := Config{&oauth2.Config{}, &FileStorage{}}
	ctx := context.Background()
	type args struct {
		ctx context.Context
		t   *oauth2.Token
	}
	tests := []struct {
		name string
		c    *Config
		args args
		want oauth2.TokenSource
	}{
		{"base-case", &conf, args{ctx, nil}, StorageTokenSource(ctx, &conf, nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.TokenSource(tt.args.ctx, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.TokenSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Client(t *testing.T) {
	conf := Config{&oauth2.Config{}, &FileStorage{}}
	ctx := context.Background()
	type args struct {
		ctx context.Context
		t   *oauth2.Token
	}
	tests := []struct {
		name string
		c    *Config
		args args
		want *http.Client
	}{
		{"base-case", &conf, args{ctx, nil}, oauth2.NewClient(ctx, conf.TokenSource(ctx, nil))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Client(tt.args.ctx, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.Client() = %v, want %v", got, tt.want)
			}
		})
	}
}
