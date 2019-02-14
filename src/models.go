package discovery

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

// Structs for the XML responses from autoconfig endpoints

// AutoConfigResponse is a parent struct used for parsing different response formats
type AutoConfigResponse struct {
	EmailProvider *EmailProvider `xml:"emailProvider"`
	Response      *Response      `xml:"Response"`
}

// EmailProvider is the domain clientconfig provider type
// Response is the other supported domain response.
type EmailProvider struct {
	ID              string           `xml:"id,attr"`
	IncomingServers []IncomingServer `xml:"incomingServer"`
	Documentation   []Documentation  `xml:"documentation"`
}

// IncomingServer contains most of the useful domain clientconfig info
type IncomingServer struct {
	Type           string `xml:"type,attr"`
	Hostname       string `xml:"hostname"`
	Port           int    `xml:"port"`
	SocketType     string `xml:"socketType"`
	Username       string `xml:"username"`
	Authentication string `xml:"authentication"`
}

// Documentation are the documentation links provided by the clientconfig info
type Documentation struct {
	URL  string              `xml:"url,attr"`
	Text []DocumentationText `xml:"descr"`
}

// DocumentationText is the localized text of a Documentation
type DocumentationText struct {
	Language string `xml:"lang,attr"`
	Text     string `xml:",chardata"`
}

// Response is the domain  AutoDiscover body.
// EmailProvider is the other supported domain response.
type Response struct {
	Account Account `xml:"Account"`
}

// Account is part of the AutoDiscover body
type Account struct {
	Protocols []Protocol `xml:"Protocol"`
}

// Protocol is where all the useful info in an AutoDiscover response lies
type Protocol struct {
	Type   string `xml:"Type"`
	Server string `xml:"Server"`
	Port   int    `xml:"Port"`
	SSL    string `xml:"SSL"`
}
