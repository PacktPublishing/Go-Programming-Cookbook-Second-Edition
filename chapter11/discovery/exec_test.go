package discovery

import (
	"errors"
	"testing"

	"github.com/hashicorp/consul/api"
)

type MockClient struct {
	MockRegister func(tags []string) error
	MockService  func(service, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error)
}

func (c *MockClient) Register(tags []string) error {
	if c.MockRegister != nil {
		return c.MockRegister(tags)
	}
	return nil
}

func (c *MockClient) Service(service, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	if c.MockService != nil {
		return c.MockService(service, tag)
	}
	return nil, nil, nil
}

func TestExec(t *testing.T) {
	type args struct {
		cli Client
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{
			&MockClient{
				MockService: func(service, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error) {
					return []*api.ServiceEntry{&api.ServiceEntry{}}, nil, nil
				},
			},
		}, false},
		{"register error", args{&MockClient{MockRegister: func(tags []string) error { return errors.New("failed") }}}, true},
		{"service error", args{
			&MockClient{
				MockService: func(service, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error) {
					return nil, nil, errors.New("failed")
				},
			},
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Exec(tt.args.cli); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
