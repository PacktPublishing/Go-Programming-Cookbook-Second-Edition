package negotiate

import (
	"net/http"
	"testing"
)

func TestGetNegotiator(t *testing.T) {
	h := make(http.Header)
	h["Content-Type"] = []string{"application/json"}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{&http.Request{Header: h}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetNegotiator(tt.args.r)
		})
	}
}
