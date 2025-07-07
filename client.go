package gohafas

// client.go deals with defining and configuring the Client

import (
	"context"
	"errors"
	"net/http"

	"github.com/xNaCly/go-hafas/language"
	"github.com/xNaCly/go-hafas/vbbraw"
)

// Client holds all necessary state and configuration for making hafas api interaction work
type Client struct {
	Baseurl             string
	Authorization       string                      // Every client using the API needs to pass a valid authentication key in every request. See 1.2.13 hafas docs
	Language            language.Language           // Language represents a language available in the HAFAS API
	ClientWithResponses *vbbraw.ClientWithResponses // Escape hatch, so the user of this api can make requests on their own
	HttpClient          vbbraw.HttpRequestDoer      // used by all methods to make HTTP requests
	Context             context.Context             // all methods pass this ctx to all HTTP requests
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
		Language:      language.EN,
		HttpClient:    &http.Client{},
		Context:       context.Background(),
	}

	for _, o := range opts {
		o(c)
	}

	return c, nil
}

// WithLanguage sets (*Client).Language = lang
func WithLanguage(lang language.Language) ClientOption {
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

// WithContext sets (*Client).Context = context
func WithContext(context context.Context) ClientOption {
	return func(c *Client) {
		c.Context = context
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
	req, err := http.NewRequestWithContext(c.Context, http.MethodGet, c.Baseurl+"/", nil)
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
