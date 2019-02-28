package webdriver

//Capabilities is a map that stores capabilities of a session.
type Capabilities map[string]interface{}

type Chrome struct {
	Args       []string `json:"args,omitempty"`
	Extensions []string `json:"extensions,omitempty"`
}

type Proxy struct {
	Type          string   `json:"proxyType,omitempty"`
	AutoconfigUrl string   `json:"proxyAutoconfigUrl,omitempty"`
	FTP           string   `json:"ftpProxy,omitempty"`
	HTTP          string   `json:"httpProxy,omitempty"`
	NoProxy       []string `json:"noProxy,omitempty"`
	SSL           string   `json:"sslProxy,omitempty"`
	Socks         string   `json:"socksProxy,omitempty"`
	SocksVersion  uint8    `json:"socksVersion,omitempty"`
	SocksUsername string   `json:"socksUsername,omitempty"`
}

func NewSocks5Proxy(host string) *Proxy {
	return &Proxy{
		Type:         "manual",
		Socks:        host,
		SocksVersion: 5,
	}
}

func (c Capabilities) Chrome(chrome *Chrome) {
	c["chromeOptions"] = chrome
}

func (c Capabilities) Proxy(proxy *Proxy) {
	c["proxy"] = proxy
}
