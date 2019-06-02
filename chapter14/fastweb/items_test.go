package main

import (
	"reflect"
	"testing"
)

func TestAddItem(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{"test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddItem(tt.args.item)
		})
	}
}

func TestReadItems(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"base-case", []string{"", "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadItems(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
