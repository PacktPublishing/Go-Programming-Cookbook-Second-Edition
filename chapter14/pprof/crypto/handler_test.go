package crypto

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGuessHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{"bad guess", args{httptest.NewRecorder(), httptest.NewRequest("GET", "/guesses?message=test", nil)}},
		{"good guess", args{httptest.NewRecorder(), httptest.NewRequest("GET", "/guesses?message=password", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GuessHandler(tt.args.w, tt.args.r)
		})
	}
}
