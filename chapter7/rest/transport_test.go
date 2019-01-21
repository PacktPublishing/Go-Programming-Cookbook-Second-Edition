package rest

import (
	"net/http"
	"testing"
)

func TestAPITransport_RoundTrip(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		t       *APITransport
		args    args
		wantErr bool
	}{
		{"base-case", &APITransport{Transport: &http.Transport{}}, args{&http.Request{Header: make(http.Header)}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.t.RoundTrip(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("APITransport.RoundTrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
