package confformat

import (
	"bytes"
	"reflect"
	"testing"
)

func TestYAMLData_ToYAML(t *testing.T) {
	type fields struct {
		Name string
		Age  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *bytes.Buffer
		wantErr bool
	}{
		{"base-case", fields{"Someone", 100}, bytes.NewBufferString("name: Someone\nage: 100\n"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := &YAMLData{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			got, err := td.ToYAML()
			if (err != nil) != tt.wantErr {
				t.Errorf("YAMLData.ToYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("YAMLData.ToYAML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYAMLData_Decode(t *testing.T) {
	type fields struct {
		Name string
		Age  int
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"base-case", fields{}, args{[]byte("name: Someone\nage: 100\n")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := &YAMLData{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if err := td.Decode(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("YAMLData.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
