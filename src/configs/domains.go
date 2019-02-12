package configs

import (
	"errors"
)

// GetKnownDomainConfig returns pre-defined configs for known domains
func GetKnownDomainConfig(username string, domain string) (*Config, error) {
	if username == "" || domain == "" {
		return nil, errors.New("missing GetKnownDomainConfig params")
	}

	// check config map
	knownConfig, found := knownDomains[domain]
	if found {
		config := *knownConfig
		// Default settings
		email := username + "@" + domain
		config.Email = email
		config.Username = email

		// Username exceptions
		if domain == "shortmail.com" || config.Provider == ProviderApple {
			config.Username = username
		}

		return &config, nil
	}

	return nil, errors.New("no config found")
}

// knownDomains are domains mapped to their base configurations
var knownDomains = map[string]*Config{
	// about.me
	"about.me": &aboutMeBaseConfig,

	// aol
	"aol.com": &aolBaseConfig,

	// gmail + google
	"gmail.com":      &gmailBaseConfig,
	"googlemail.com": &gmailBaseConfig,
	"google.com":     &googleMailBaseConfig,

	// Apple
	"mac.com":    &appleBaseConfig,
	"me.com":     &appleBaseConfig,
	"icloud.com": &appleBaseConfig,

	// Microsoft
	"charter.com":       &microsoftBaseConfig,
	"compaq.net":        &microsoftBaseConfig,
	"hotmail.be":        &microsoftBaseConfig,
	"hotmail.ca":        &microsoftBaseConfig,
	"hotmail.ch":        &microsoftBaseConfig,
	"hotmail.cl":        &microsoftBaseConfig,
	"hotmail.co.id":     &microsoftBaseConfig,
	"hotmail.co.il":     &microsoftBaseConfig,
	"hotmail.co.in":     &microsoftBaseConfig,
	"hotmail.co.jp":     &microsoftBaseConfig,
	"hotmail.co.kr":     &microsoftBaseConfig,
	"hotmail.co.nz":     &microsoftBaseConfig,
	"hotmail.co.th":     &microsoftBaseConfig,
	"hotmail.co.uk":     &microsoftBaseConfig,
	"hotmail.co.za":     &microsoftBaseConfig,
	"hotmail.com":       &microsoftBaseConfig,
	"hotmail.com.ar":    &microsoftBaseConfig,
	"hotmail.com.au":    &microsoftBaseConfig,
	"hotmail.com.br":    &microsoftBaseConfig,
	"hotmail.com.hk":    &microsoftBaseConfig,
	"hotmail.com.tr":    &microsoftBaseConfig,
	"hotmail.com.tw":    &microsoftBaseConfig,
	"hotmail.com.vn":    &microsoftBaseConfig,
	"hotmail.cz":        &microsoftBaseConfig,
	"hotmail.de":        &microsoftBaseConfig,
	"hotmail.dk":        &microsoftBaseConfig,
	"hotmail.es":        &microsoftBaseConfig,
	"hotmail.fi":        &microsoftBaseConfig,
	"hotmail.fr":        &microsoftBaseConfig,
	"hotmail.gr":        &microsoftBaseConfig,
	"hotmail.hu":        &microsoftBaseConfig,
	"hotmail.it":        &microsoftBaseConfig,
	"hotmail.lt":        &microsoftBaseConfig,
	"hotmail.lv":        &microsoftBaseConfig,
	"hotmail.my":        &microsoftBaseConfig,
	"hotmail.nl":        &microsoftBaseConfig,
	"hotmail.no":        &microsoftBaseConfig,
	"hotmail.ph":        &microsoftBaseConfig,
	"hotmail.rs":        &microsoftBaseConfig,
	"hotmail.se":        &microsoftBaseConfig,
	"hotmail.sg":        &microsoftBaseConfig,
	"hotmail.sk":        &microsoftBaseConfig,
	"live.at":           &microsoftBaseConfig,
	"live.be":           &microsoftBaseConfig,
	"live.ca":           &microsoftBaseConfig,
	"live.cl":           &microsoftBaseConfig,
	"live.cn":           &microsoftBaseConfig,
	"live.co.in":        &microsoftBaseConfig,
	"live.co.kr":        &microsoftBaseConfig,
	"live.co.uk":        &microsoftBaseConfig,
	"live.co.za":        &microsoftBaseConfig,
	"live.com":          &microsoftBaseConfig,
	"live.com.ar":       &microsoftBaseConfig,
	"live.com.au":       &microsoftBaseConfig,
	"live.com.co":       &microsoftBaseConfig,
	"live.com.mx":       &microsoftBaseConfig,
	"live.com.my":       &microsoftBaseConfig,
	"live.com.pe":       &microsoftBaseConfig,
	"live.com.ph":       &microsoftBaseConfig,
	"live.com.pk":       &microsoftBaseConfig,
	"live.com.pt":       &microsoftBaseConfig,
	"live.com.sg":       &microsoftBaseConfig,
	"live.com.ve":       &microsoftBaseConfig,
	"live.de":           &microsoftBaseConfig,
	"live.dk":           &microsoftBaseConfig,
	"live.fi":           &microsoftBaseConfig,
	"live.fr":           &microsoftBaseConfig,
	"live.hk":           &microsoftBaseConfig,
	"live.ie":           &microsoftBaseConfig,
	"live.in":           &microsoftBaseConfig,
	"live.it":           &microsoftBaseConfig,
	"live.jp":           &microsoftBaseConfig,
	"live.ma":           &microsoftBaseConfig,
	"live.nl":           &microsoftBaseConfig,
	"live.no":           &microsoftBaseConfig,
	"live.ph":           &microsoftBaseConfig,
	"live.ru":           &microsoftBaseConfig,
	"live.se":           &microsoftBaseConfig,
	"livemail.com.br":   &microsoftBaseConfig,
	"livemail.tw":       &microsoftBaseConfig,
	"messengeruser.com": &microsoftBaseConfig,
	"msn.com":           &microsoftBaseConfig,
	"outlook.at":        &microsoftBaseConfig,
	"outlook.be":        &microsoftBaseConfig,
	"outlook.cl":        &microsoftBaseConfig,
	"outlook.co.id":     &microsoftBaseConfig,
	"outlook.co.il":     &microsoftBaseConfig,
	"outlook.co.nz":     &microsoftBaseConfig,
	"outlook.co.th":     &microsoftBaseConfig,
	"outlook.com":       &microsoftBaseConfig,
	"outlook.com.ar":    &microsoftBaseConfig,
	"outlook.com.au":    &microsoftBaseConfig,
	"outlook.com.br":    &microsoftBaseConfig,
	"outlook.com.gr":    &microsoftBaseConfig,
	"outlook.com.pe":    &microsoftBaseConfig,
	"outlook.com.tr":    &microsoftBaseConfig,
	"outlook.com.vn":    &microsoftBaseConfig,
	"outlook.cz":        &microsoftBaseConfig,
	"outlook.de":        &microsoftBaseConfig,
	"outlook.dk":        &microsoftBaseConfig,
	"outlook.es":        &microsoftBaseConfig,
	"outlook.fr":        &microsoftBaseConfig,
	"outlook.hu":        &microsoftBaseConfig,
	"outlook.ie":        &microsoftBaseConfig,
	"outlook.in":        &microsoftBaseConfig,
	"outlook.it":        &microsoftBaseConfig,
	"outlook.jp":        &microsoftBaseConfig,
	"outlook.kr":        &microsoftBaseConfig,
	"outlook.lv":        &microsoftBaseConfig,
	"outlook.my":        &microsoftBaseConfig,
	"outlook.ph":        &microsoftBaseConfig,
	"outlook.pt":        &microsoftBaseConfig,
	"outlook.sa":        &microsoftBaseConfig,
	"outlook.sg":        &microsoftBaseConfig,
	"outlook.sk":        &microsoftBaseConfig,
	"passport.com":      &microsoftBaseConfig,
	"passport.net":      &microsoftBaseConfig,
	"vcache.com":        &microsoftBaseConfig,
	"webtv.net":         &microsoftBaseConfig,
	"windowslive.com":   &microsoftBaseConfig,
	"windowslive.es":    &microsoftBaseConfig,

	// Shortmail
	"shortmail.com": &ShortmailBaseConfig,

	// Yahoo
	"yahoo.com":      &yahooBaseConfig,
	"ca.yahoo.com":   &yahooBaseConfig,
	"qc.yahoo.com":   &yahooBaseConfig,
	"yahoo.at":       &yahooBaseConfig,
	"yahoo.be":       &yahooBaseConfig,
	"yahoo.ca":       &yahooBaseConfig,
	"yahoo.co.in":    &yahooBaseConfig,
	"yahoo.co.id":    &yahooBaseConfig,
	"yahoo.co.il":    &yahooBaseConfig,
	"yahoo.co.jp":    &yahooBaseConfig,
	"yahoo.co.nz":    &yahooBaseConfig,
	"yahoo.co.th":    &yahooBaseConfig,
	"yahoo.co.uk":    &yahooBaseConfig,
	"yahoo.co.za":    &yahooBaseConfig,
	"yahoo.com.ar":   &yahooBaseConfig,
	"yahoo.com.au":   &yahooBaseConfig,
	"yahoo.com.br":   &yahooBaseConfig,
	"yahoo.com.co":   &yahooBaseConfig,
	"yahoo.com.hr":   &yahooBaseConfig,
	"yahoo.com.hk":   &yahooBaseConfig,
	"yahoo.com.my":   &yahooBaseConfig,
	"yahoo.com.mx":   &yahooBaseConfig,
	"yahoo.com.ph":   &yahooBaseConfig,
	"yahoo.com.sg":   &yahooBaseConfig,
	"yahoo.com.tw":   &yahooBaseConfig,
	"yahoo.com.tr":   &yahooBaseConfig,
	"yahoo.com.vn":   &yahooBaseConfig,
	"yahoo.ae":       &yahooBaseConfig,
	"yahoo.ch":       &yahooBaseConfig,
	"yahoo.cz":       &yahooBaseConfig,
	"yahoo.de":       &yahooBaseConfig,
	"yahoo.dk":       &yahooBaseConfig,
	"yahoo.es":       &yahooBaseConfig,
	"yahoo.fi":       &yahooBaseConfig,
	"yahoo.fr":       &yahooBaseConfig,
	"yahoo.gr":       &yahooBaseConfig,
	"yahoo.hu":       &yahooBaseConfig,
	"yahoo.ie":       &yahooBaseConfig,
	"yahoo.it":       &yahooBaseConfig,
	"yahoo.nl":       &yahooBaseConfig,
	"yahoo.no":       &yahooBaseConfig,
	"yahoo.pl":       &yahooBaseConfig,
	"yahoo.pt":       &yahooBaseConfig,
	"yahoo.ro":       &yahooBaseConfig,
	"yahoo.ru":       &yahooBaseConfig,
	"yahoo.se":       &yahooBaseConfig,
	"rocketmail.com": &yahooBaseConfig,
	"ymail.com":      &yahooBaseConfig,
	"att.net":        &yahooBaseConfig,
}
