package discovery

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
)

// Tests the mozilla autoconfig function + parsing
func TestGetMozillaAutoConfig(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// load test data
	res, err := ioutil.ReadFile("testdata/mozilla_200.xml")
	if err != nil {
		t.Error(err)
	}

	url := "https://autoconfig.thunderbird.net/v1.1/comcast.net"
	username := "test"
	domain := "comcast.net"

	httpmock.RegisterResponder("GET", url,
		httpmock.NewBytesResponder(200, res))

	config, err := GetAutoConfig(url, username, domain)
	if err != nil {
		t.Error(err)
	}

	if config.Provider != ProviderGeneric {
		t.Error("Expected type GENERIC got " + string(config.Provider))
	}

	if config.Server != "imap.comcast.net" {
		t.Error("Expected imap.comcast.net got " + config.Server)
	}

	if len(config.Documentation) != 1 {
		t.Errorf("Expected documentation length to be 1 got %v", len(config.Documentation))
	}
}

// Tests the domain autodiscover function + parsing
func TestGetAutoDiscover(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// load test data
	res, err := ioutil.ReadFile("testdata/autodiscover_200.xml")
	if err != nil {
		t.Error(err)
	}

	url := "http://autoconfig.comcast.net/mail/config-v1.1.xml?emailaddress=test@comcast.net"
	username := "test"
	domain := "comcast.net"

	httpmock.RegisterResponder("GET", url,
		httpmock.NewBytesResponder(200, res))

	config, err := GetAutoConfig(url, username, domain)
	if err != nil {
		t.Error(err)
	}

	if config.Provider != ProviderGeneric {
		t.Error("Expected type GENERIC got " + string(config.Provider))
	}

	if config.Server != "imap.comcast.net" {
		t.Error("Expected imap.comcast.net got " + config.Server)
	}
}

// Tests the autoconfig function with no domain - should fail
func TestGetAutoConfigNoDomain(t *testing.T) {

	_, err := GetAutoConfig("", "", "")
	if err == nil {
		t.Error("no domain, should have errored")
	}
}

// Tests the autoconfig function with a non-200 returned - should fail
func TestGetAutoConfigNon200(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://autoconfig.thunderbird.net/v1.1/one.com"
	username := "test"
	domain := "one.com"

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(404, ``))

	_, err := GetAutoConfig(url, username, domain)
	if err == nil {
		t.Error("404, should have errored")
	}
	if err.Error() != "unexpected response: 404" {
		t.Error(err)
	}
}

// Tests the autoconfig function with an empty - should fail
func TesGetAutoConfigEmptyBody(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://autoconfig.thunderbird.net/v1.1/one.com"
	username := "test"
	domain := "one.com"

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, ``))

	_, err := GetAutoConfig(url, username, domain)
	if err == nil {
		t.Error("empty body, should have errored")
	}
	if err.Error() != "EOF" {
		t.Error(err)
	}
}

// Tests the autoconfig function with no imap incomingServer - should fail
func TestGetAutoConfigNoImapServer(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// load test data
	res, err := ioutil.ReadFile("testdata/verizon.xml")
	if err != nil {
		t.Error(err)
	}

	url := "https://autoconfig.thunderbird.net/v1.1/verizon.net"
	username := "test"
	domain := "verizon.com"

	httpmock.RegisterResponder("GET", url,
		httpmock.NewBytesResponder(200, res))

	_, err = GetAutoConfig(url, username, domain)
	if err == nil {
		t.Error("missing incoming server, should have errored")
	}
}

// Tests the autoconfig function with an unsupported xml response format - should fail
func TestGetAutoConfigUnsupportedFormat(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://autoconfig.thunderbird.net/v1.1/gmail.com"
	username := "test"
	domain := "one.com"

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, `<dumbFormat id='dealWithIt'></dumbFormat>`))

	_, err := GetAutoConfig(url, username, domain)
	if err == nil {
		t.Error("unsupported format, should have errored")
	}
	if err.Error() != "domain config response did not match a know format" {
		t.Error(err)
	}
}

// Tests the autoconfig function with a broken xml response - should fail
func TestGetAutoConfigMalformedBody(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://autoconfig.thunderbird.net/v1.1/gmail.com"
	username := "test"
	domain := "one.com"

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, `<clientConfig version="1.1">`))

	_, err := GetAutoConfig(url, username, domain)
	if err == nil {
		t.Error("bad xml, should have errored")
	}
	if err.Error() != "XML syntax error on line 1: unexpected EOF" {
		t.Error(err)
	}
}
