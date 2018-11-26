package currency

import "testing"

func TestConvertStringDollarsToPennies(t *testing.T) {
	type args struct {
		amount string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		//"0.01", "0.10", "1", "1.1", "1.11", "1.111", "50", "50.49", "50.119"
		{"1 cent", args{"0.01"}, 1, false},
		{"-1 cent", args{"-0.01"}, -1, false},
		{"10 cents", args{"0.10"}, 10, false},
		{"1 dollar", args{"1.00"}, 100, false},
		{"1 dollar no decimal", args{"1"}, 100, false},
		{"-1 dollar", args{"-1.00"}, -100, false},
		{"1 dollar 11 cents", args{"1.11"}, 111, false},
		{"invalid decimal 1", args{"1.111"}, 0, true},
		{"invalid decimal 2", args{"1.1"}, 0, true},
		{"invalid character", args{"a"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertStringDollarsToPennies(tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertStringDollarsToPennies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertStringDollarsToPennies() = %v, want %v", got, tt.want)
			}
		})
	}
}
