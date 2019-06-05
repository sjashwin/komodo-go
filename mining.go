package komodo

import "net/http"

func (c *Client) GetBlockSubsidy(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getblocksubsidy", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetBlockTemplate(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getblocktemplate", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetLocalSolPS() (*http.Response, error) {
	if body, err := c.generateCurl("getlocalsolps", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetMiningInfo() (*http.Response, error) {
	if body, err := c.generateCurl("verifytxoutproof", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetNetworkHashPS(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getnetworkhashps", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetNetworkSolPS(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getnetworksolps", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) PrioritiseTransaction(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("prioritisetransaction", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) SubmitBlock(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("submitblock", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
