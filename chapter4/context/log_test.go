package context

import (
	"context"
	"reflect"
	"testing"

	"github.com/apex/log"
)

func Test_getFields(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *log.Fields
	}{
		{"base-case", args{context.Background()}, &log.Fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFields(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
