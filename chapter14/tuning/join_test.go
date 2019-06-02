package tuning

import "testing"

func Test_join(t *testing.T) {
	type args struct {
		vals []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"base-case", args{[]string{"a", "b", "c"}}, "a b c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := join(tt.args.vals...); got != tt.want {
				t.Errorf("join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_join(b *testing.B) {
	b.Run("one", func(b *testing.B) {
		one := []string{"1"}
		for i := 0; i < b.N; i++ {
			join(one...)
		}
	})
	b.Run("five", func(b *testing.B) {
		five := []string{"1", "2", "3", "4", "5"}
		for i := 0; i < b.N; i++ {
			join(five...)
		}
	})

	b.Run("ten", func(b *testing.B) {
		ten := []string{"1", "2", "3", "4", "5",
			"6", "7", "8", "9", "10"}
		for i := 0; i < b.N; i++ {
			join(ten...)
		}
	})
}
