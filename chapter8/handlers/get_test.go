package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		code int
	}{
		{"base-case", args{httptest.NewRecorder(), httptest.NewRequest("GET", "/name?name=test", nil)}, http.StatusOK},
		{"bad method", args{httptest.NewRecorder(), httptest.NewRequest("POST", "/name?name=test", nil)}, http.StatusMethodNotAllowed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HelloHandler(tt.args.w, tt.args.r)
			w := tt.args.w.(*httptest.ResponseRecorder)
			if got, want := w.Code, tt.code; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}
