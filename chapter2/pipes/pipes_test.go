package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	f1, err := ioutil.TempFile(".", "tmp")
	if err != nil {
		t.Error(err)
	}
	defer f1.Close()
	defer os.Remove(f1.Name())

	f1.Write([]byte("test case"))
	f1.Seek(0, 0)

	type args struct {
		f *os.File
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"base-case", args{f1}, map[string]int{"test": 1, "case": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordCount(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
