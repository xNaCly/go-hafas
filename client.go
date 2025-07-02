package gohafas

// client.go deals with defining and configuring Client

import (
	"errors"

	"github.com/xNaCly/go-hafas/vbbraw"
)

// Client holds the Baseurl of the hafas remote, the Authentication token and
// the internal vbbraw.ClientWithResponses for raw hafas usage by the consumer
type Client struct {
	Baseurl             string
	Authentication      string   // Every client using the API needs to pass a valid authentication key in every request. - 1.2.13 Authentication from HAFAS.api.pdf
	Language            Language // Language represents a language available in the HAFAS API
	ClientWithResponses vbbraw.ClientWithResponses
}

// ClientOption defines a function that configures the Client.
type ClientOption func(*Client)

func NewClient(baseurl, authentication string, opts ...ClientOption) (*Client, error) {
	if baseurl == "" {
		return nil, errors.New("`baseurl` cannot be empty")
	}
	if authentication == "" {
		return nil, errors.New("`authentication` token cannot be empty")
	}
	c := &Client{
		Baseurl:        baseurl,
		Authentication: authentication,
		Language:       EN,
	}

	for _, o := range opts {
		o(c)
	}

	return c, nil
}

// WithLanguage sets (*Client).Language = lang
func WithLanguage(lang Language) ClientOption {
	return func(c *Client) {
		c.Language = lang
	}
}

var unimplemented = errors.New("Unimplemented")

// Init establishes a connection to the hafas remote and initializes internal state
func (c *Client) Init() error { return unimplemented }

// Ping attempts to ping the remote and returns an error if said ping fails
func (c *Client) Ping() error { return unimplemented }
