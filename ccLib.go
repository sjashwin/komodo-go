package komodo

import "net/http"

// GetInfo returns an object containing various states info.
// Returns http response and error
func (c *Client) GetInfo() (*http.Response, error) {
	if body, err := c.generateCurl("getinfo", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
