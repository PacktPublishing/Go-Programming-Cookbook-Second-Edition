package async

import (
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		client     *http.Client
		bufferSize int
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{http.DefaultClient, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewClient(tt.args.client, tt.args.bufferSize)
		})
	}
}

func TestClient_AsyncGet(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		c    *Client
		args args
	}{
		{"base-case", NewClient(http.DefaultClient, 10), args{"https://www.google.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.AsyncGet(tt.args.url)
		})
	}
}
