package oauthstore

import (
	"context"
	"reflect"
	"testing"

	"golang.org/x/oauth2"
)

func Test_storageTokenSource_Token(t *testing.T) {
	conf := Config{&oauth2.Config{}, &FileStorage{}}
	ctx := context.Background()
	s := storageTokenSource{&conf, conf.Config.TokenSource(ctx, nil)}
	tests := []struct {
		name    string
		s       *storageTokenSource
		wantErr bool
	}{
		{"base-case", &s, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s.Token()
			if (err != nil) != tt.wantErr {
				t.Errorf("storageTokenSource.Token() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestStorageTokenSource(t *testing.T) {
	conf := Config{&oauth2.Config{}, &FileStorage{}}
	ctx := context.Background()
	ts := conf.Config.TokenSource(ctx, nil)
	sts := &storageTokenSource{&conf, ts}
	type args struct {
		ctx context.Context
		c   *Config
		t   *oauth2.Token
	}
	tests := []struct {
		name string
		args args
		want oauth2.TokenSource
	}{
		{"base-case", args{ctx, &conf, nil}, sts},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StorageTokenSource(tt.args.ctx, tt.args.c, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageTokenSource() = %v, want %v", got, tt.want)
			}
		})
	}
}
