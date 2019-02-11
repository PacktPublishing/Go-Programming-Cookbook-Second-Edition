package controllers

import "testing"

func TestMemStorage_Get(t *testing.T) {
	tests := []struct {
		name string
		m    *MemStorage
		want string
	}{
		{"base-case", &MemStorage{value: "123"}, "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Get(); got != tt.want {
				t.Errorf("MemStorage.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemStorage_Put(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		m    *MemStorage
		args args
	}{
		{"base-case", &MemStorage{}, args{"value"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Put(tt.args.s)
		})
	}
}
