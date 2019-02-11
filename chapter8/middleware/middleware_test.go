package middleware

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestApplyMiddleware(t *testing.T) {
	h := func(http.ResponseWriter, *http.Request) {}
	type args struct {
		h          http.HandlerFunc
		middleware []Middleware
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{h, []Middleware{SetID(1)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ApplyMiddleware(tt.args.h, tt.args.middleware...)
		})
	}
}

func TestLogger(t *testing.T) {
	h := func(http.ResponseWriter, *http.Request) {}
	type args struct {
		l *log.Logger
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{log.New(os.Stdout, "", 0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Logger(tt.args.l)(h)(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
		})
	}
}
