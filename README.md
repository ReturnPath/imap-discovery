# IMAP Discovery

A library to help 'discover' IMAP connection parameters.
IMAP is hard, and it's hard to know how to connect to a given email service provider.

Some providers, like Gmail, have well-known connection parameters.
Others may choose to publish connection parameters via self-hosted XML documents, or publishing them to [Mozilla Thunderbird Autoconfiguration project](https://developer.mozilla.org/en-US/docs/Mozilla/Thunderbird/Autoconfiguration). There may also be MX records to check!

This library checks all these sources and either gives you all the successful results, or the one that we believe is the most likely to succeed.

Our recommended order, and the one checked by the main library entrypoint, `DiscoverImapConfig`, is:

- Known Domains
- Autoconfig documents hosted by the domain
- Mozilla's autoconfig service
- MX records for the domain

## Usage

```go
// Get one config
config, err := discovery.DiscoverImapConfig("email@example.com")

// Check all possible options
configs, err := discovery.DiscoverAllImapConfigs("email@example.com")

// Check a specific method
// note the email split in the specific methods
config, err = GetMozillaAutoConfig("email", "example.com)
```

## Response shape

```go
// Config represents the IMAP Config response from this service
type Config struct {
	Email         string              `json:"email"`
	Username      string              `json:"username"`
	Provider      string              `json:"provider"`
	Server        string              `json:"server"`
	Port          int                 `json:"port"`
	SSL           bool                `json:"ssl"`
	OAuth         bool                `json:"oauth"`
	Documentation []ImapDocumentation `json:"documentation"`
}

// ImapDocumentation represents provided documentation
type ImapDocumentation struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}
```
