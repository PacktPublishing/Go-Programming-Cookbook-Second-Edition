package dns

import (
	"testing"
)

func TestLookupAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base case", args{"www.google.com"}, false},
		{"case 2", args{"A"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LookupAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookupAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLookup_String(t *testing.T) {
	type fields struct {
		cname string
		hosts []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{"", []string{}}, ""},
		{"not empty", fields{"cname", []string{"golang.org"}}, "cname IN A golang.org\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Lookup{
				cname: tt.fields.cname,
				hosts: tt.fields.hosts,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("Lookup.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
