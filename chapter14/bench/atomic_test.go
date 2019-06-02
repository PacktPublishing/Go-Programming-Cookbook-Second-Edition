package bench

import "testing"

func TestAtomicCounter_Add(t *testing.T) {
	type fields struct {
		value int64
	}
	type args struct {
		amount int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base-case", fields{0}, args{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AtomicCounter{
				value: tt.fields.value,
			}
			c.Add(tt.args.amount)
		})
	}
}

func TestAtomicCounter_Read(t *testing.T) {
	type fields struct {
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{"ten", fields{10}, 10},
		{"-nine", fields{-9}, -9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AtomicCounter{
				value: tt.fields.value,
			}
			if got := c.Read(); got != tt.want {
				t.Errorf("AtomicCounter.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAtomicCounterAdd(b *testing.B) {
	c := AtomicCounter{0}
	for n := 0; n < b.N; n++ {
		c.Add(1)
	}
}

func BenchmarkAtomicCounterRead(b *testing.B) {
	c := AtomicCounter{0}
	for n := 0; n < b.N; n++ {
		c.Read()
	}
}

func BenchmarkAtomicCounterAddRead(b *testing.B) {
	c := AtomicCounter{0}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Add(1)
			c.Read()
		}
	})
}
