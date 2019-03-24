package waitgroup

import (
	"errors"
	"testing"
)

func TestGetURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{"https://www.google.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCrawlError_Add(t *testing.T) {
	type fields struct {
		Errors []string
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base-case", fields{make([]string, 0)}, args{errors.New("failed")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CrawlError{
				Errors: tt.fields.Errors,
			}
			c.Add(tt.args.err)
		})
	}
}

func TestCrawlError_Error(t *testing.T) {
	type fields struct {
		Errors []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"no errors", fields{[]string{}}, "All Errors: "},
		{"one error", fields{[]string{"error1", "error2"}}, "All Errors: error1,error2"},
		{"two errors", fields{[]string{"error1"}}, "All Errors: error1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CrawlError{
				Errors: tt.fields.Errors,
			}
			if got := c.Error(); got != tt.want {
				t.Errorf("CrawlError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrawlError_Valid(t *testing.T) {
	type fields struct {
		Errors []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CrawlError{
				Errors: tt.fields.Errors,
			}
			if got := c.Valid(); got != tt.want {
				t.Errorf("CrawlError.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
