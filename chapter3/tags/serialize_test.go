package tags

import "testing"

func TestSerializeStructStrings(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"base-case", args{Person{Name: "test"}}, "name:test;city:;State:;", false},
		{"base-case", args{Person{Name: "name", City: "city", State: "state", Misc: "misc", Year: 2017}}, "name:name;city:city;State:state;", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SerializeStructStrings(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("SerializeStructStrings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SerializeStructStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
