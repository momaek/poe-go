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
	Headers http.Header

	Proxy string // TODO
}

func (c *Config) HTTPHeaders() http.Header {
	headers := http.Header{}
	for k, h := range c.Headers {
		if len(h) > 0 {
			headers.Add(k, h[0])
		}
	}

	for k, h := range basicHeader {
		headers.Set(k, h)
	}

	cookie := &http.Cookie{
		Name:   "p-b",
		Value:  c.Token,
		Domain: "poe.com",
	}
	headers.Add("Cookie", cookie.String())

	return headers
}

type Transport struct {
	conf *Config
	tr   http.RoundTripper
}

var defaultHeaders = http.Header{
	"User-Agent":                []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"},
	"Accept":                    []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"},
	"Accept-Encoding":           []string{"gzip, deflate, br"},
	"Accept-Language":           []string{"en-US,en;q=0.5"},
	"Te":                        []string{"trailers"},
	"Upgrade-Insecure-Requests": []string{"1"},
}

var basicHeader = map[string]string{
	"Referrer":       "https://poe.com/",
	"Origin":         "https://poe.com",
	"Host":           "poe.com",
	"Sec-Fetch-Dest": "empty",
	"Sec-Fetch-Mode": "cors",
	"Sec-Fetch-Site": "same-origin",
}

// RoundTrip implement http.RoundTripper
func (t *Transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	headers := t.conf.HTTPHeaders()
	for k, h := range headers {
		for _, v := range h {
			req.Header.Add(k, v)
		}
	}

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
