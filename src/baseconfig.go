package discovery

// Basic configs for known domains

// genericBaseConfig is a base config for generic imap servers
var genericBaseConfig = Config{
	Provider: ProviderGeneric,
	OAuth:    false,
	SSL:      true,
	Port:     993,
}

// aboutMeBaseConfig is the known base config for about.me
var aboutMeBaseConfig = Config{
	Provider: ProviderGeneric,
	Server:   "imap.about.me",
	OAuth:    false,
	SSL:      true,
	Port:     993,
}

// aolBaseConfig is the known base config for aol
var aolBaseConfig = Config{
	Provider: ProviderAOL,
	Server:   "imap.aol.com",
	OAuth:    false,
	SSL:      true,
	Port:     993,
}

// gmailBaseConfig is the known base config for gmail
var gmailBaseConfig = Config{
	Provider:      ProviderGMail,
	Server:        "imap.gmail.com",
	OAuth:         true,
	SSL:           true,
	Port:          993,
	Documentation: googleDocumentation,
}

// googleMailBaseConfig is the known base config for google apps mail
var googleMailBaseConfig = Config{
	Provider:      ProviderGoogleapps,
	Server:        "imap.googlemail.com",
	OAuth:         true,
	SSL:           true,
	Port:          993,
	Documentation: googleDocumentation,
}

// GoogleDocumentation are some hard-coded documentation links for gmail accounts
var googleDocumentation = []ImapDocumentation{
	ImapDocumentation{
		Description: "How to enable IMAP/POP3 in GMail",
		URL:         "http://mail.google.com/support/bin/answer.py?answer=13273",
	},
	ImapDocumentation{
		Description: "How to configure email clients for IMAP",
		URL:         "http://mail.google.com/support/bin/topic.py?topic=12806",
	},
}

// appleBaseConfig is the known base config for me.com, icloud.com, and mac.com
var appleBaseConfig = Config{
	Provider:      ProviderApple,
	Server:        "mail.me.com",
	OAuth:         false,
	SSL:           true,
	Port:          993,
	Documentation: AppleDocumentation,
}

// AppleDocumentation are some hard-coded documentation links for apple accounts
var AppleDocumentation = []ImapDocumentation{
	ImapDocumentation{
		Description: "Using app-specific passwords",
		URL:         "https://support.apple.com/en-us/HT204397",
	},
}

// microsoftBaseConfig is the known base config for microsoft
var microsoftBaseConfig = Config{
	Provider: ProviderMicrosoft,
	Server:   "imap-mail.outlook.com",
	OAuth:    true,
	SSL:      true,
	Port:     993,
}

// ShortmailBaseConfig is the known base config for shortmail
var shortmailBaseConfig = Config{
	Provider: ProviderGeneric,
	Server:   "imap.shortmail.com",
	OAuth:    false,
	SSL:      true,
	Port:     993,
}

// yahooBaseConfig is the known base config for yahoo
var yahooBaseConfig = Config{
	Provider: ProviderYahoo,
	Server:   "imap.mail.yahoo.com",
	OAuth:    false,
	SSL:      true,
	Port:     993,
}
