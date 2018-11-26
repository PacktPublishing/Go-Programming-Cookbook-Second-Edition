package collections

import (
	"reflect"
	"testing"
)

func TestLowerCaseData(t *testing.T) {
	type args struct {
		w WorkWith
	}
	tests := []struct {
		name string
		args args
		want WorkWith
	}{
		{"base-case", args{WorkWith{"Test", 0}}, WorkWith{"test", 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerCaseData(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LowerCaseData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIncrementVersion(t *testing.T) {
	type args struct {
		w WorkWith
	}
	tests := []struct {
		name string
		args args
		want WorkWith
	}{
		{"base-case", args{WorkWith{"Test", 0}}, WorkWith{"Test", 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IncrementVersion(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncrementVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOldVersion(t *testing.T) {
	type args struct {
		v  int
		ws WorkWith
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"base-case", args{1, WorkWith{"Test", 0}}, false},
		{"base-case", args{1, WorkWith{"Test", 1}}, true},
		{"base-case", args{1, WorkWith{"Test", 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OldVersion(tt.args.v)(tt.args.ws); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OldVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
