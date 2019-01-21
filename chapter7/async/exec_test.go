package async

import (
	"net/http"
	"testing"
)

func TestFetchAll(t *testing.T) {
	type args struct {
		urls []string
		c    *Client
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{nil, NewClient(http.DefaultClient, 10)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FetchAll(tt.args.urls, tt.args.c)
		})
	}
}
