package decorator

import (
	"net/http"
	"reflect"
	"testing"
)

func TestTransportFunc_RoundTrip(t *testing.T) {
	req, _ := http.NewRequest("get", "www.google.com", nil)
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		tf      TransportFunc
		args    args
		wantErr bool
	}{
		{
			"base-case",
			TransportFunc(func(*http.Request) (*http.Response, error) { return nil, nil }),
			args{req},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.tf.RoundTrip(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransportFunc.RoundTrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDecorate(t *testing.T) {
	type args struct {
		t   http.RoundTripper
		rts []Decorator
	}
	tests := []struct {
		name string
		args args
		want http.RoundTripper
	}{
		{"base-case", args{&http.Transport{}, nil}, &http.Transport{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decorate(tt.args.t, tt.args.rts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decorate() = %v, want %v", got, tt.want)
			}
		})
	}
}
