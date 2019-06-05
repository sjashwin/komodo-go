package komodo

import "net/http"

func (c *Client) AddNode(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("addnode", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ClearBanned() (*http.Response, error) {
	if body, err := c.generateCurl("clearbanned", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) DisconnectNode(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("disconnectnode", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetAddedNodeInfo(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getaddednodeinfo", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetConnectionCount(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getconnectioncount", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetDeprecationInfo() (*http.Response, error) {
	if body, err := c.generateCurl("getdeprecationinfo", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetNetTotals() (*http.Response, error) {
	if body, err := c.generateCurl("getnettotals", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetNetworkInfo() (*http.Response, error) {
	if body, err := c.generateCurl("getnetworkinfo", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetPeerInfo() (*http.Response, error) {
	if body, err := c.generateCurl("getpeerinfo", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ListBanned() (*http.Response, error) {
	if body, err := c.generateCurl("listbanned", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) Ping() (*http.Response, error) {
	if body, err := c.generateCurl("ping", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) SetBan(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("setban", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
