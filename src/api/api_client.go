/*
Copyright © 2020 Mehdi Bounya <mehdi.bounya@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
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