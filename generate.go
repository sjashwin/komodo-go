package komodo

import "net/http"

func (c *Client) GetGenerate() (*http.Response, error) {
	if body, err := c.generateCurl("getgenerate", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) SetGenerate(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("setgenerate", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
