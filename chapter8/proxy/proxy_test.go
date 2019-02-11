package proxy

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxy_ServeHTTP(t *testing.T) {
	type fields struct {
		Client  *http.Client
		BaseURL string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base-case", fields{http.DefaultClient, ""}, args{httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Proxy{
				Client:  tt.fields.Client,
				BaseURL: tt.fields.BaseURL,
			}
			p.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}
