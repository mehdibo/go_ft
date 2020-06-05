package api

import (
	"fmt"
	"net/http"
	"time"
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
	var sleep time.Duration = 5

	// TODO: add max retries
	for {
		resp, err := c.authenticatedClient.Do(req)
		if resp.StatusCode != http.StatusTooManyRequests {
			return resp, err
		}
		fmt.Printf("The API is complaining about too many requests, pausing for %d seconds\n", sleep)
		time.Sleep(sleep)
		sleep += sleep * (20/100)
	}

}

func (c *Client) Get(url string) (resp *http.Response, err error)  {
	req, err := http.NewRequest("GET", c.apiEndpoint+url, nil)

	if err != nil {
		return nil, err
	}

	return c.do(req)
}