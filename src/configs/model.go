package configs

const (
	// ProviderGeneric is for generic IMAP providers
	ProviderGeneric = "GENERIC"
	// ProviderGMail is for GMail IMAP
	ProviderGMail = "GMAIL"
	// ProviderGoogleapps is for google apps (g-suite) IMAP
	ProviderGoogleapps = "GOOGLEAPPS"
	// ProviderMicrosoft is for Microsoft IMAP
	ProviderMicrosoft = "MICROSOFT"
	// ProviderYahoo is for Yahoo IMAP
	ProviderYahoo = "YAHOO"
	// ProviderAOL is for AOL IMAP
	ProviderAOL = "AOL"
	// ProviderApple is for Apple IMAP
	ProviderApple = "APPLE"
)

// Config represents the IMAP Config response from this service
type Config struct {
	Email         string
	Username      string
	Provider      string
	Server        string
	Port          int
	SSL           bool
	OAuth         bool
	Documentation []ImapDocumentation
}

// ImapDocumentation represents provided documentation
type ImapDocumentation struct {
	URL         string
	Description string
}
