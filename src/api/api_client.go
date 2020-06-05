package api

import (
	"net/http"
)

type Client struct {
	apiEndpoint string
	authenticatedClient *http.Client
}

func Create(apiEndpoint string, authenticatedClient *http.Client) *Client  {
	return &Client{
		apiEndpoint:         apiEndpoint,
		authenticatedClient: authenticatedClient,
	}
}

func (c *Client) do(req *http.Request) (*http.Response, error)  {
	return c.authenticatedClient.Do(req)
}

func (c *Client) Get(url string) (resp *http.Response, err error)  {
	req, err := http.NewRequest("GET", c.apiEndpoint+url, nil)

	if err != nil {
		return nil, err
	}

	return c.do(req)
}