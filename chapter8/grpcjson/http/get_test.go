package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter8/grpcjson/internal"
	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter8/grpcjson/keyvalue"
)

func TestController_GetHandler(t *testing.T) {
	k := internal.NewKeyValue()
	k.Set(context.Background(), &keyvalue.SetKeyValueRequest{Key: "test", Value: "value"})
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		c    *Controller
		args args
	}{
		{"no key", &Controller{KeyValue: internal.NewKeyValue()}, args{httptest.NewRecorder(), httptest.NewRequest("GET", "/?key=test", nil)}},
		{"key", &Controller{KeyValue: k}, args{httptest.NewRecorder(), httptest.NewRequest("GET", "/?key=test", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.GetHandler(tt.args.w, tt.args.r)
		})
	}
}
