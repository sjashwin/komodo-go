package komodo

import "net/http"

func (c *Client) CreateRawTransaction(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("createRawTransaction", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) DecodeRawTransaction(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("decoderawtransaction", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) DecodeScript(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("decodescript", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) FundRawTransaction(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("fundrawtransaction", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetRawTransaction(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getrawtransaction", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) SendRawTransaction(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("sendrawtransaction", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) SignRawTransaction(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("signrawtransaction", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
