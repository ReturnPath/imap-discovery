package discovery

import (
	"errors"
	"strings"

	"github.com/ReturnPath/contextio-discovery/src/configs"
)

// DiscoverImapConfig is the library entrypoint!
// The options that are more likely to succeed are tried first.
// In order: Known Domain Configs, Domain Autoconfig endpoints, Mozilla Autoconfig information, and finally MX records for the domain itself
func DiscoverImapConfig(email string) (*configs.Config, error) {

	// split email into username and domain
	emailParts := strings.Split(email, "@")
	if len(emailParts) != 2 {
		return nil, errors.New("invalid email address")
	}

	username := emailParts[0]
	domain := emailParts[1]

	// check known domains
	config, err := configs.GetKnownDomainConfig(username, domain)
	if err == nil {
		return config, nil
	}

	// check for autoconfig from the domain
	config, err = configs.GetDomainAutoConfig(username, domain)
	if err == nil {
		return config, nil
	}

	// check mozilla's config
	config, err = configs.GetMozillaAutoConfig(username, domain)
	if err == nil {
		return config, nil
	}

	// check the MX records for the domain
	config, err = configs.GetMXRecord(domain, email)
	if err == nil {
		return config, nil
	}

	// No results :(
	return nil, errors.New("unable to discover configuration")
}
