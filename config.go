package aimlapi

import (
	"net/http"
)

const (
	aimlAPIURLv1                   = "https://api.aimlapi.com/v1"
	defaultEmptyMessagesLimit uint = 300
)

// ClientConfig is a configuration of a client.
// authToken is your API key
type ClientConfig struct {
	authToken          string
	XTitle             string
	HttpReferer        string
	BaseURL            string
	HTTPClient         *http.Client
	EmptyMessagesLimit uint
}

func DefaultConfig(auth, xTitle, httpReferer string) (ClientConfig, error) {
	return ClientConfig{
		authToken:          auth,
		XTitle:             xTitle,
		HttpReferer:        httpReferer,
		HTTPClient:         &http.Client{},
		BaseURL:            aimlAPIURLv1,
		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}, nil
}

func (c ClientConfig) WithHttpClientConfig(client *http.Client) ClientConfig {
	c.HTTPClient = client
	return c
}
