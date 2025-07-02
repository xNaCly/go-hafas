package gohafas

// client.go deals with defining and configuring the Client

import (
	"context"
	"errors"
	"net/http"

	"github.com/xNaCly/go-hafas/vbbraw"
)

// Client holds the Baseurl of the hafas remote, the Authentication token and
// the internal vbbraw.ClientWithResponses for raw hafas usage by the consumer
type Client struct {
	Baseurl             string
	Authorization       string   // Every client using the API needs to pass a valid authentication key in every request. - 1.2.13 Authentication from HAFAS.api.pdf
	Language            Language // Language represents a language available in the HAFAS API
	ClientWithResponses *vbbraw.ClientWithResponses
	HttpClient          vbbraw.HttpRequestDoer
}

// ClientOption defines a function that configures the Client.
type ClientOption func(*Client)

func NewClient(baseurl, authorization string, opts ...ClientOption) (*Client, error) {
	if baseurl == "" {
		return nil, errors.New("`baseurl` cannot be empty")
	}
	if authorization == "" {
		return nil, errors.New("`authorization` cannot be empty")
	}
	c := &Client{
		Baseurl:       baseurl,
		Authorization: authorization,
		Language:      EN,
		HttpClient:    &http.Client{},
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

// WithHttpClient sets (*Client).HttpClient = client
func WithHttpClient(client vbbraw.HttpRequestDoer) ClientOption {
	return func(c *Client) {
		c.HttpClient = client
	}
}

// Init establishes a connection to the hafas remote and initializes internal
// state, must be called before any methods are called
func (c *Client) Init() (err error) {
	c.ClientWithResponses, err = vbbraw.NewClientWithResponses(c.Baseurl, func(cwr *vbbraw.Client) error {
		cwr.Client = c.HttpClient
		return nil
	}, vbbraw.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		// probably not that good to do it on every request like this, but at this point, who cares
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Authorization", "Bearer "+c.Authorization)

		q := req.URL.Query()
		q.Set("lang", c.Language)
		req.URL.RawQuery = q.Encode()

		return nil
	}))

	if err != nil {
		return errors.Join(errors.New("Failed to initialize gohafas.Client"), err)
	}
	return nil
}

// Ping attempts to ping the remote and returns an error if said ping fails
func (c *Client) Ping() error {
	req, err := http.NewRequest(http.MethodGet, c.Baseurl+"/", nil)
	if err != nil {
		return errors.Join(errors.New("Failed to create the ping request"), err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.Authorization)

	res, err := c.HttpClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return errors.Join(errors.New("Failed to do the ping request"), err)
	}

	return nil
}
