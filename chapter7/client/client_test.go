package client

import (
	"net/http"
	"reflect"
	"testing"
)

func TestSetup(t *testing.T) {
	type args struct {
		isSecure bool
		nop      bool
	}
	tests := []struct {
		name string
		args args
		want *http.Client
	}{
		{"insecure noops", args{false, false}, http.DefaultClient},
		{"insecure, ops", args{false, true}, http.DefaultClient},
		{"secure, ops", args{true, false}, http.DefaultClient},
		{"secure, noops", args{true, true}, http.DefaultClient},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Setup(tt.args.isSecure, tt.args.nop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Setup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNopTransport_RoundTrip(t *testing.T) {
	type args struct {
		in0 *http.Request
	}
	tests := []struct {
		name    string
		n       *NopTransport
		args    args
		want    *http.Response
		wantErr bool
	}{
		{"base-case", &NopTransport{}, args{&http.Request{}}, &http.Response{StatusCode: http.StatusTeapot}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NopTransport{}
			got, err := n.RoundTrip(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("NopTransport.RoundTrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NopTransport.RoundTrip() = %v, want %v", got, tt.want)
			}
		})
	}
}
