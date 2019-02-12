package configs

import (
	"errors"
	"net"
	"strings"
)

// GetMXRecord checks the MX records of a given domain and returns a config for known records
func GetMXRecord(domain string, email string) (*Config, error) {
	if domain == "" || email == "" {
		return nil, errors.New("missing GetMXRecord params")
	}

	// Query MX records
	records, err := net.LookupMX(domain)
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, errors.New("no records found for domain " + domain)
	}

	// default config
	config := genericBaseConfig
	config.Email = email
	config.Username = email

	for i := 0; i < len(records); i++ {
		host := strings.ToLower(records[i].Host)

		// check for google apps records
		if strings.Contains(host, "aspmx.l.google.com") {
			config = googleMailBaseConfig
			config.Email = email
			config.Username = email
			return &config, nil
		}

		// check for rackspace records
		if strings.Contains(host, "emailsrvr.com") {
			config.Server = "secure.emailsrvr.com"
			return &config, nil
		}

		// check for liveoffice records
		if strings.Contains(host, "hex.namehub.com") {
			config.Server = "exchange.liveoffice.com"
			return &config, nil
		}
	}

	return nil, errors.New("no usable records for domain " + domain + "")
}
