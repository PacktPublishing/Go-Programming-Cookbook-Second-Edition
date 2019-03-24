package atomic

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewSafeMap(t *testing.T) {
	tests := []struct {
		name string
		want SafeMap
	}{
		{"base-case", SafeMap{m: make(map[string]string), mu: &sync.RWMutex{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSafeMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSafeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeMap_Set(t *testing.T) {
	type fields struct {
		m  map[string]string
		mu *sync.RWMutex
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base-case", fields{make(map[string]string), &sync.RWMutex{}}, args{"key", "value"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SafeMap{
				m:  tt.fields.m,
				mu: tt.fields.mu,
			}
			s.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestSafeMap_Get(t *testing.T) {
	type fields struct {
		m  map[string]string
		mu *sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"key found", fields{make(map[string]string), &sync.RWMutex{}}, args{"key1"}, "value1", false},
		{"key missing", fields{make(map[string]string), &sync.RWMutex{}}, args{"key2"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SafeMap{
				m:  tt.fields.m,
				mu: tt.fields.mu,
			}
			s.Set("key1", "value1")
			got, err := s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("SafeMap.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SafeMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
