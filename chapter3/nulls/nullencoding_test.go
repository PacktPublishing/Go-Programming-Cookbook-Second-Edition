package nulls

import "testing"

/*func Test_nullInt_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		v       nullInt64
		want    []byte
		wantErr bool
	}{
		{"1 valid", nullInt64{Int64: 1, Valid: true}, []byte("1"), false},
		{"1 invalid", nullInt64{Int64: 1, Valid: false}, []byte("null"), false},
		{"0 valid", nullInt64{Int64: 0, Valid: true}, []byte("0"), false},
		{"0 invalid", nullInt64{Int64: 0, Valid: false}, []byte("null"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("nullInt.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nullInt.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

func TestNullEncoding(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NullEncoding(); (err != nil) != tt.wantErr {
				t.Errorf("NullEncoding() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
