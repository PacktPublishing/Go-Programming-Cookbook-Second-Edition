package waitgroup

import (
	"reflect"
	"testing"
)

func TestCrawl(t *testing.T) {
	type args struct {
		sites []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"base-case", args{[]string{"https://www.google.com"}}, []int{200}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Crawl(tt.args.sites)
			if (err != nil) != tt.wantErr {
				t.Errorf("Crawl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Crawl() = %v, want %v", got, tt.want)
			}
		})
	}
}
