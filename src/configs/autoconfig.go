package configs

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// DomainFormat is the generic domain autoconfig url format
const DomainFormat = "http://autoconfig.%s/mail/config-v1.1.xml?emailaddress=%s"

// MozillaFormat is the mozilla endpoint for domain config
const MozillaFormat = "https://autoconfig.thunderbird.net/v1.1/%s"

// GetDomainAutoConfig fetches the domain config information.
func GetDomainAutoConfig(username string, domain string) (*Config, error) {
	if username == "" || domain == "" {
		return nil, errors.New("missing GetDomainAutoConfig params")
	}

	url := fmt.Sprintf(DomainFormat, domain, username+"@"+domain)
	return GetAutoConfig(url, username, domain)
}

// GetMozillaAutoConfig fetches the mozilla default domain config information.
// https://developer.mozilla.org/en/Thunderbird/Autoconfiguration
func GetMozillaAutoConfig(username string, domain string) (*Config, error) {
	if username == "" || domain == "" {
		return nil, errors.New("missing GetMozillaAutoConfig params")
	}

	url := fmt.Sprintf(MozillaFormat, domain)
	return GetAutoConfig(url, username, domain)
}

// GetAutoConfig fetches the xml autoconfig info from a given url.
// It attempts to parse one of two supported formats, Autodiscover and clientConfig.
func GetAutoConfig(url string, username string, domain string) (*Config, error) {
	if url == "" || username == "" || domain == "" {
		return nil, errors.New("missing GetAutoConfig params")
	}

	// Make the domain config request
	config, err := makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// There are two supported response types, so parse it and see what's there
	var resp AutoConfigResponse
	err = xml.Unmarshal(config, &resp)
	if err != nil {
		return nil, err
	}

	if resp.EmailProvider != nil {
		return ConvertClientConfigResponse(resp.EmailProvider, username, domain)
	}

	if resp.Response != nil {
		return ConvertAutoDiscoverResponse(resp.Response, username, domain)
	}

	return nil, errors.New("domain config response did not match a know format")
}

// ConvertClientConfigResponse parses the client config response into an ImapConfig.
func ConvertClientConfigResponse(config *EmailProvider, username string, domain string) (*Config, error) {
	var imapServer *IncomingServer

	// make sure there's an IMAP config present
	for i := 0; i < len(config.IncomingServers); i++ {
		server := config.IncomingServers[i]
		if server.Type == "imap" {
			imapServer = &server
			break
		}
	}

	if imapServer == nil {
		return nil, errors.New("no IMAP config found")
	}

	// simple fields
	imapConfig := Config{
		Email:    username + "@" + domain,
		Provider: ProviderGeneric,
		Port:     imapServer.Port,
		Server:   imapServer.Hostname,
		SSL:      imapServer.SocketType == "SSL",
		OAuth:    imapServer.Authentication == "OAuth2",
	}

	// username
	switch imapServer.Username {
	case "%EMAILADDRESS%":
		imapConfig.Username = username + "@" + domain
	case "%EMAILLOCALPART%":
		imapConfig.Username = username
	default:
		imapConfig.Username = imapServer.Username
	}

	// documentation
	for i := 0; i < len(config.Documentation); i++ {
		doc := config.Documentation[i]
		text := ""

		// documentation text
		for k := 0; k < len(doc.Text); k++ {
			text = text + doc.Text[k].Text + " "
		}

		pbDoc := ImapDocumentation{
			Description: strings.TrimSpace(text),
			URL:         doc.URL,
		}

		imapConfig.Documentation = append(imapConfig.Documentation, pbDoc)
	}

	return &imapConfig, nil
}

// ConvertAutoDiscoverResponse parses the autodiscover response into an ImapConfig.
func ConvertAutoDiscoverResponse(config *Response, username string, domain string) (*Config, error) {
	imapProtocol := &Protocol{}

	// make sure there's an IMAP protocol present
	for i := 0; i < len(config.Account.Protocols); i++ {
		protocol := config.Account.Protocols[i]
		if protocol.Type == "IMAP" {
			imapProtocol = &protocol
			break
		}
	}

	if imapProtocol == nil {
		return nil, errors.New("no IMAP protocol found")
	}

	// simple fields
	imapConfig := Config{
		Email:    username + "@" + domain,
		Username: username + "@" + domain,
		Provider: ProviderGeneric,
		Port:     imapProtocol.Port,
		Server:   imapProtocol.Server,
		SSL:      imapProtocol.SSL == "on",
		OAuth:    false,
	}

	return &imapConfig, nil
}

// makeHTTPRequest is just an http helper func
func makeHTTPRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response: " + res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
