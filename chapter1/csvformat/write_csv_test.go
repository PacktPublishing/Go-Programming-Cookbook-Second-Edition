package csvformat

import (
	"bytes"
	"reflect"
	"testing"
)

func TestBooks_ToCSV(t *testing.T) {
	tests := []struct {
		name    string
		books   *Books
		wantW   string
		wantErr bool
	}{
		{
			name: "base-case",
			books: &Books{
				Book{
					Author: "F Scott Fitzgerald",
					Title:  "The Great Gatsby",
				},
			},
			wantW:   "Author,Title\nF Scott Fitzgerald,The Great Gatsby\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := tt.books.ToCSV(w); (err != nil) != tt.wantErr {
				t.Errorf("Books.ToCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Books.ToCSV() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestWriteCSVOutput(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteCSVOutput(); (err != nil) != tt.wantErr {
				t.Errorf("WriteCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteCSVBuffer(t *testing.T) {
	tests := []struct {
		name    string
		want    *bytes.Buffer
		wantErr bool
	}{
		{"base-case", bytes.NewBufferString("Author,Title\nF Scott Fitzgerald,The Great Gatsby\nJ D Salinger,The Catcher in the Rye\n"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteCSVBuffer()
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteCSVBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WriteCSVBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}
