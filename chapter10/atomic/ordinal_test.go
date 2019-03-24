package atomic

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewOrdinal(t *testing.T) {
	tests := []struct {
		name string
		want *Ordinal
	}{
		{"base-case", &Ordinal{once: &sync.Once{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrdinal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrdinal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdinal_Init(t *testing.T) {
	type fields struct {
		ordinal uint64
		once    *sync.Once
	}
	type args struct {
		val uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base-case", fields{0, &sync.Once{}}, args{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Ordinal{
				ordinal: tt.fields.ordinal,
				once:    tt.fields.once,
			}
			o.Init(tt.args.val)
		})
	}
}

func TestOrdinal_GetOrdinal(t *testing.T) {
	type fields struct {
		ordinal uint64
		once    *sync.Once
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{"base-case", fields{0, &sync.Once{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Ordinal{
				ordinal: tt.fields.ordinal,
				once:    tt.fields.once,
			}
			if got := o.GetOrdinal(); got != tt.want {
				t.Errorf("Ordinal.GetOrdinal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdinal_Increment(t *testing.T) {
	type fields struct {
		ordinal uint64
		once    *sync.Once
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"base-case", fields{0, &sync.Once{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Ordinal{
				ordinal: tt.fields.ordinal,
				once:    tt.fields.once,
			}
			o.Increment()
		})
	}
}
