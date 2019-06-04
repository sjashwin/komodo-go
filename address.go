package komodo

import (
	"net/http"
)

// GetAddressBalance returns the balance in the current address
func (c *Client) GetAddressBalance(param Params) (*http.Response, error) {
	if body, err := c.generateCurl("getaddressbalance", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetAddressDeltas(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getaddressbalance", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetAddressMempool(address string) (*http.Response, error) {
	if body, err := c.generateCurl("getaddressmempool", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetAddressTXIDS(address string, params ...Params) (*http.Response, error) {
	if body, err := c.generateCurl("getaddresstxids", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetAddressUTXOS(address string, params ...Params) (*http.Response, error) {
	if body, err := c.generateCurl("getaddressutxos", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// TODO: needs to be tested. No address is passed
func (c *Client) GetSnapShot(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getsnapshot", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
