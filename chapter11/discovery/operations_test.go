package discovery

import (
	"testing"

	"github.com/hashicorp/consul/api"
)

func Test_client_Register(t *testing.T) {

	type args struct {
		tags []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(api.DefaultConfig(), "test", "test", 123)
			if err != nil {
				t.Errorf("failed to make new client")
			}
			if err := c.Register(tt.args.tags); (err != nil) != tt.wantErr {
				t.Errorf("client.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_Service(t *testing.T) {

	type args struct {
		service string
		tag     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{"test", "test"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(api.DefaultConfig(), "test", "test", 123)
			if err != nil {
				t.Errorf("failed to make new client")
			}
			_, _, err = c.Service(tt.args.service, tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Service() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
