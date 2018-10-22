package csvformat

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestReadCSV(t *testing.T) {
	type args struct {
		b io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []Movie
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "base-case",
			args: args{
				bytes.NewBufferString(`
movie title;director;year released
Guardians of the Galaxy Vol. 2;James Gunn;2017
`),
			},
			want:    []Movie{Movie{"Guardians of the Galaxy Vol. 2", "James Gunn", 2017}},
			wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadCSV(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddMoviesFromText(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddMoviesFromText(); (err != nil) != tt.wantErr {
				t.Errorf("AddMoviesFromText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
