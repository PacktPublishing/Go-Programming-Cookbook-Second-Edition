package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreetingHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		code int
	}{
		{"base-case", args{httptest.NewRecorder(), httptest.NewRequest("POST", "/greeting", nil)}, http.StatusOK},
		{"bad method", args{httptest.NewRecorder(), httptest.NewRequest("GET", "/greeting", nil)}, http.StatusMethodNotAllowed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GreetingHandler(tt.args.w, tt.args.r)
			w := tt.args.w.(*httptest.ResponseRecorder)
			if got, want := w.Code, tt.code; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}
