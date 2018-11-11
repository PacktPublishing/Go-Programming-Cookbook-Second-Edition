package confformat

import (
	"bytes"
	"reflect"
	"testing"
)

func TestJSONData_ToJSON(t *testing.T) {
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
		{"base-case", fields{"Someone", 100}, bytes.NewBufferString(`{"name":"Someone","age":100}`), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := &JSONData{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			got, err := td.ToJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONData.ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONData.ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONData_Decode(t *testing.T) {
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
		{"base-case", fields{}, args{[]byte(`{"name":"Someone","age":100}`)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := &JSONData{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if err := td.Decode(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("JSONData.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOtherJSONExamples(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OtherJSONExamples(); (err != nil) != tt.wantErr {
				t.Errorf("OtherJSONExamples() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
