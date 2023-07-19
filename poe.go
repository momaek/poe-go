package poego

// Client poego client
// do every request
type Client struct {
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

// NewClient new poe-go client
func NewClient(conf *Config) *Client {
	return &Client{conf}
}

// SendMessage send message to a bot
func (c *Client) SendMessage() {}
