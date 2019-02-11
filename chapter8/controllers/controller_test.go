package controllers

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	storage := MemStorage{}
	type args struct {
		storage Storage
	}
	tests := []struct {
		name string
		args args
		want *Controller
	}{
		{"base-case", args{storage: &storage}, &Controller{storage: &storage}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
