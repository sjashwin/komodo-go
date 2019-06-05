package komodo

import "net/http"

func (c *Client) CoinSupply(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("coinsupply", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetBlockHash() (*http.Response, error) {
	if body, err := c.generateCurl("getblockhash", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetBlock(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getblock", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetBlockchainInfo(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getblockchaininfo", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetBlockCount() (*http.Response, error) {
	if body, err := c.generateCurl("getblockcount", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetBlockHashes(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getblockhashes", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetBlockHeader(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getblockheader", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetChainTips() (*http.Response, error) {
	if body, err := c.generateCurl("getchaintips", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetChainTXStats() (*http.Response, error) {
	if body, err := c.generateCurl("getchaintxstats", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetDifficulty(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getdifficulty", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetLastSegIDStakes(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getlastsegidstakes", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetMemPoolinfo(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getmempoolinfo", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetRawMemPool(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getrawmempool", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetSpentInfo(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getspentinfo", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetTXOut(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("gettxout", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetTXOutProof(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("gettxoutproof", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetTXOutSetInfo() (*http.Response, error) {
	if body, err := c.generateCurl("gettxoutsetinfo", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) KVSearch(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("kvsearch", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) KVUpdate(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("kvupdate", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) MinerIDs(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("minerids", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) Notaries(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("notaires", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) VerifyChain(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("verifychain", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) VerifyTXOutProof(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("verifytxoutproof", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
