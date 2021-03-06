package config

import (
	"fmt"

	"bosh-dns/dns/manager"
)

const loopbackAddress = "127.0.0.1"

func NewRecursorReader(dnsManager manager.DNSManager, dnsNameServers []string) recursorReader {
	return recursorReader{
		manager:        dnsManager,
		dnsNameServers: dnsNameServers,
	}
}

type recursorReader struct {
	manager        manager.DNSManager
	dnsNameServers []string
}

func (r recursorReader) Get() ([]string, error) {
	recursors := []string{}

	nameservers, err := r.manager.Read()
	if err != nil {
		return nil, err
	}

	for _, server := range nameservers {
		if r.isValid(server) {
			recursors = append(recursors, fmt.Sprintf("%s:53", server))
		}
	}

	return recursors, nil
}

func (r recursorReader) isNameServer(s string) bool {
	for _, nameServer := range r.dnsNameServers {
		if nameServer == s {
			return true
		}
	}

	return false
}

func (r recursorReader) isValid(server string) bool {
	return !r.isNameServer(server) &&
		server != loopbackAddress &&
		server != ""
}
