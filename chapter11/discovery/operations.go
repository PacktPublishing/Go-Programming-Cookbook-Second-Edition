package discovery

import "github.com/hashicorp/consul/api"

// Register adds our service to consul
func (c *client) Register(tags []string) error {
	reg := &api.AgentServiceRegistration{
		ID:      c.name,
		Name:    c.name,
		Port:    c.port,
		Address: c.address,
		Tags:    tags,
	}
	return c.client.Agent().ServiceRegister(reg)
}

// Service return a service
func (c *client) Service(service, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	return c.client.Health().Service(service, tag, false, nil)
}
