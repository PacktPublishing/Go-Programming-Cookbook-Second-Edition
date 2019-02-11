package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetID(t *testing.T) {
	h := func(http.ResponseWriter, *http.Request) {}
	type args struct {
		start int64
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetID(tt.args.start)(h)(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
		})
	}
}

func TestGetID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"base-case", args{context.WithValue(context.Background(), ID, "123")}, "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetID(tt.args.ctx); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}
