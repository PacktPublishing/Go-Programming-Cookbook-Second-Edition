package main

import (
	"testing"

	"github.com/valyala/fasthttp"
)

func TestGetItems(t *testing.T) {
	type args struct {
		ctx *fasthttp.RequestCtx
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{&fasthttp.RequestCtx{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetItems(tt.args.ctx)
		})
	}
}

func TestAddItems(t *testing.T) {
	type args struct {
		ctx *fasthttp.RequestCtx
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{&fasthttp.RequestCtx{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddItems(tt.args.ctx)
		})
	}
}
