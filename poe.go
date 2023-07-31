package poego

import "net/http"

// Client poego client
// do every request
type Client struct {
	*http.Client
	conf *Config
}

// Config poe-go config struct
type Config struct {
	// Log into Poe on any desktop web browser, then open your browser's developer tools (also known as "inspect") and look for the value of the p-b cookie in the following menus:
	// Chromium: Devtools > Application > Cookies > poe.com
	// Firefox: Devtools > Storage > Cookies
	// Safari: Devtools > Storage > Cookies
	Token string
	// The device ID to use. If this is not specified, it will be randomly generated and stored on the disk.
	DeviceID string

	// The headers to use. This defaults to the headers specified in
	// poego.Headers
	Headers []string

	Proxy string // TODO
}

type Transport struct {
	conf *Config
	tr   http.RoundTripper
}

// RoundTrip implement http.RoundTripper
func (t *Transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	return t.tr.RoundTrip(req)
}

// NewClient new poe-go client
func New(conf *Config) *Client {
	return &Client{
		Client: &http.Client{Transport: newTransport(conf, http.DefaultTransport)},
		conf:   conf,
	}
}

func NewWithHTTPClient(conf *Config, cli *http.Client) *Client {
	c := &Client{conf: conf}

	// wrap transport
	tr := newTransport(conf, cli.Transport)
	cli.Transport = tr

	c.Client = cli
	return c
}

func newTransport(conf *Config, tr http.RoundTripper) *Transport {
	return &Transport{
		conf: conf,
		tr:   tr,
	}
}
