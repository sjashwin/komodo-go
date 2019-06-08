package komodo

import "encoding/json"

func (c *Client) CoinSupply(params Params) (*CoinSupply, error) {
	data, err := c.send("coinsupply", params)
	if err != nil {
		return nil, err
	}
	var v = &CoinSupply{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetBestBlockHash(params Params) (*GetBestBlockHash, error) {
	data, err := c.send("getbestblockhash", params)
	if err != nil {
		return nil, err
	}
	var v = &GetBestBlockHash{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetBlock(params Params) (*GetBlock, error) {
	data, err := c.send("getblock", params)
	if err != nil {
		return nil, err
	}
	var v = &GetBlock{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetBlockchainInfo(params Params) (*GetBlockchainInfo, error) {
	data, err := c.send("getblockchaininfo", params)
	if err != nil {
		return nil, err
	}
	var v = &GetBlockchainInfo{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetBlockCount(params Params) (*GetBlockCount, error) {
	data, err := c.send("getblockcount", params)
	if err != nil {
		return nil, err
	}
	var v = &GetBlockCount{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetBlockHash(params Params) (*GetBlockHash, error) {
	data, err := c.send("getblockhash", params)
	if err != nil {
		return nil, err
	}
	var v = &GetBlockHash{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetBlockHashes(params Params) (*GetBlockHashes, error) {
	data, err := c.send("getblockhashes", params)
	if err != nil {
		return nil, err
	}
	var v = &GetBlockHashes{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetChainTips(params Params) (*GetChainTips, error) {
	data, err := c.send("getchaintips", params)
	if err != nil {
		return nil, err
	}
	var v = &GetChainTips{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetChainTXStats(params Params) (*GetChainTXStats, error) {
	data, err := c.send("getchaintxstats", params)
	if err != nil {
		return nil, err
	}
	var v = &GetChainTXStats{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetDifficulty(params Params) (*GetDifficulty, error) {
	data, err := c.send("getchaintxstats", params)
	if err != nil {
		return nil, err
	}
	var v = &GetDifficulty{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetLastSegIDStakes(params Params) (*GetLastSegIDStakes, error) {
	data, err := c.send("getlastsegidstakes", params)
	if err != nil {
		return nil, err
	}
	var v = &GetLastSegIDStakes{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetMemPoolInfo(params Params) (*GetMemPoolInfo, error) {
	data, err := c.send("getmempoolinfo", params)
	if err != nil {
		return nil, err
	}
	var v = &GetMemPoolInfo{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetSpentInfo(params Params) (*GetSpentInfo, error) {
	data, err := c.send("getspentinfo", params)
	if err != nil {
		return nil, err
	}
	var v = &GetSpentInfo{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetTXOut(params Params) (*GetTXOut, error) {
	data, err := c.send("gettxout", params)
	if err != nil {
		return nil, err
	}
	var v = &GetTXOut{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetTXOutProof(params Params) (*GetTXOutProof, error) {
	data, err := c.send("gettxoutproof", params)
	if err != nil {
		return nil, err
	}
	var v = &GetTXOutProof{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetTXOutSetInfo(params Params) (*GetTXOutSetInfo, error) {
	data, err := c.send("gettxoutsetinfo", params)
	if err != nil {
		return nil, err
	}
	var v = &GetTXOutSetInfo{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) KVSearch(params Params) (*KVSearch, error) {
	data, err := c.send("kvsearch", params)
	if err != nil {
		return nil, err
	}
	var v = &KVSearch{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) KVUpdate(params Params) (*KVUpdate, error) {
	data, err := c.send("kvupdate", params)
	if err != nil {
		return nil, err
	}
	var v = &KVUpdate{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) MinerIDs(params Params) (*MinerIDs, error) {
	data, err := c.send("minerids", params)
	if err != nil {
		return nil, err
	}
	var v = &MinerIDs{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) Notaries(params Params) (*Notary, error) {
	data, err := c.send("notaries", params)
	if err != nil {
		return nil, err
	}
	var v = &Notary{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) VerifyChain(params Params) (*VerifyChain, error) {
	data, err := c.send("verifychain", params)
	if err != nil {
		return nil, err
	}
	var v = &VerifyChain{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) VerifyTXOutProof(params Params) (*VerifyTXOutProof, error) {
	data, err := c.send("gettxoutproof", params)
	if err != nil {
		return nil, err
	}
	var v = &VerifyTXOutProof{}
	return v, json.Unmarshal(data, v)
}
