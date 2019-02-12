package configs

import (
	"testing"
)

// Tests the known domains function
func TestGetKnownDomainConfig(t *testing.T) {

	config, err := GetKnownDomainConfig("test", "gmail.com")
	if err != nil {
		t.Error("should have found gmail account")
	}

	if config.Username != "test@gmail.com" {
		t.Error("incorrect formatting of username")
	}
}

// Tests the known domains function with no email or domain - should fail
func TestGetKnownDomainConfigMissingParams(t *testing.T) {

	_, err := GetKnownDomainConfig("", "")
	if err == nil {
		t.Error("missing params, should have errored")
	}
}

// Tests the known domains function with an unknown domain- should fail
func TestGetKnownDomainConfigUnknownDomain(t *testing.T) {

	_, err := GetKnownDomainConfig("bruce", "bruceiscool.codes")
	if err == nil {
		t.Error("unknown domain, should have errored")
	}
}
