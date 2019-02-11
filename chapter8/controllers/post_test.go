package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_SetValue(t *testing.T) {
	c := New(&MemStorage{})
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		c    *Controller
		args args
	}{
		{"base-case", c, args{httptest.NewRecorder(), httptest.NewRequest("POST", "/", nil)}},
		{"bad method", c, args{httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.SetValue(tt.args.w, tt.args.r)
		})
	}
}
