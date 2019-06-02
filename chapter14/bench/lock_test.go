package bench

import (
	"sync"
	"testing"
)

func TestCounter_Add(t *testing.T) {
	type fields struct {
		value int64
		mu    *sync.RWMutex
	}
	type args struct {
		amount int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base-case", fields{0, &sync.RWMutex{}}, args{10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Counter{
				value: tt.fields.value,
				mu:    tt.fields.mu,
			}
			c.Add(tt.args.amount)
		})
	}
}

func TestCounter_Read(t *testing.T) {
	type fields struct {
		value int64
		mu    *sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{"ten", fields{10, &sync.RWMutex{}}, 10},
		{"-nine", fields{-9, &sync.RWMutex{}}, -9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Counter{
				value: tt.fields.value,
				mu:    tt.fields.mu,
			}
			if got := c.Read(); got != tt.want {
				t.Errorf("Counter.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCounterAdd(b *testing.B) {
	c := Counter{0, &sync.RWMutex{}}
	for n := 0; n < b.N; n++ {
		c.Add(1)
	}
}

func BenchmarkCounterRead(b *testing.B) {
	c := Counter{0, &sync.RWMutex{}}
	for n := 0; n < b.N; n++ {
		c.Read()
	}
}

func BenchmarkCounterAddRead(b *testing.B) {
	c := Counter{0, &sync.RWMutex{}}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Add(1)
			c.Read()
		}
	})
}
