package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_GetValue(t *testing.T) {
	c := New(&MemStorage{})
	type args struct {
		UseDefault bool
		w          http.ResponseWriter
		r          *http.Request
	}
	tests := []struct {
		name string
		c    *Controller
		args args
	}{
		{"base-case", c, args{true, httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil)}},
		{"don't use default", c, args{false, httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil)}},
		{"bad method", c, args{true, httptest.NewRecorder(), httptest.NewRequest("POST", "/", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.GetValue(tt.args.UseDefault)(tt.args.w, tt.args.r)
		})
	}
}
