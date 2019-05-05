package discovery

import (
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestNewClient(t *testing.T) {
	type args struct {
		config  *api.Config
		address string
		name    string
		port    int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{api.DefaultConfig(), "test", "test", 123}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.args.config, tt.args.address, tt.args.name, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
