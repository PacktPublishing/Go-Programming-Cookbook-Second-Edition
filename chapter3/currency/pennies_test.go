package currency

import "testing"

func TestConvertPenniesToDollarString(t *testing.T) {
	type args struct {
		amount int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1 cent", args{1}, "0.01"},
		{"10 cents", args{10}, "0.10"},
		{"1 dollar", args{100}, "1.00"},
		{"1 dollar 11 cents", args{111}, "1.11"},
		{"-9 dollar 53 cents", args{-953}, "-9.53"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertPenniesToDollarString(tt.args.amount); got != tt.want {
				t.Errorf("ConvertPenniestoDollarString() = %v, want %v", got, tt.want)
			}
		})
	}
}
