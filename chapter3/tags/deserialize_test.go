package tags

import "testing"

func TestDeSerializeStructStrings(t *testing.T) {
	p := Person{}
	type args struct {
		s   string
		res interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{"name:name;city:city;State:state;", &p}, false},
		{"base-case", args{"name:name;city:city;State:state;", p}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeSerializeStructStrings(tt.args.s, tt.args.res); (err != nil) != tt.wantErr {
				t.Errorf("DeSerializeStructStrings() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
