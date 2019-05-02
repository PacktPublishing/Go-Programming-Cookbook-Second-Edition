package dns

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
)

// Lookup holds the DNS information we care about
type Lookup struct {
	cname string
	hosts []string
}

// We can use this to print the lookup object
func (d *Lookup) String() string {
	result := ""
	for _, host := range d.hosts {
		result += fmt.Sprintf("%s IN A %s\n", d.cname, host)
	}
	return result
}

// LookupAddress returns a DNSLookup consisting of a cname and host
// for a given address
func LookupAddress(address string) (*Lookup, error) {
	cname, err := net.LookupCNAME(address)
	if err != nil {
		return nil, errors.Wrap(err, "error looking up CNAME")
	}
	hosts, err := net.LookupHost(address)
	if err != nil {
		return nil, errors.Wrap(err, "error looking up HOST")
	}

	return &Lookup{cname: cname, hosts: hosts}, nil
}
