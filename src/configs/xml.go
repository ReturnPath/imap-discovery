package configs

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
