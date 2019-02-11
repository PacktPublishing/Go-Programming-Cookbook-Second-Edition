package proxy

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxy_ProcessRequest(t *testing.T) {
	type fields struct {
		Client  *http.Client
		BaseURL string
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"base-case", fields{http.DefaultClient, ""}, args{httptest.NewRequest("GET", "/", nil)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Proxy{
				Client:  tt.fields.Client,
				BaseURL: tt.fields.BaseURL,
			}
			if err := p.ProcessRequest(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ProcessRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopyResponse(t *testing.T) {
	header := make(http.Header)
	header.Add("test", "value")
	type args struct {
		w    http.ResponseWriter
		resp *http.Response
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{httptest.NewRecorder(), &http.Response{Header: header, Body: ioutil.NopCloser(bytes.NewBufferString("test"))}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CopyResponse(tt.args.w, tt.args.resp)
		})
	}
}
