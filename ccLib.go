package komodo

import "net/http"

func (c *Client) GetInfo() (*http.Response, error) {
	if body, err := c.generateCurl("getinfo", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
