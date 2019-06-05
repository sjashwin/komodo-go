package komodo

import "net/http"

func (c *Client) CreateMultiSig(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("createmultisig", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) DecodeCCOpRet(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("decodeccoperet", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) EstimateFee(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("estimatefee", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) EstimatePriority(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("estimatepriority", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) InvalidateBlock(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("invalidateblock", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ReconsiderBlock(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("reconsiderblock", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) TXNotarizedConfirmed(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("txnotarizedconfirmed", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ValidateAddress(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("validateaddress", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) VerfiyAddress(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("verifyaddress", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
