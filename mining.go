package komodo

import "encoding/json"

func (c *Client) GetBlockSubsidy(params Params) (*GetBlockSubsidy, error) {
	data, err := c.send("getblocksubsidy", params)
	if err != nil {
		return nil, err
	}
	var v = &GetBlockSubsidy{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetBlockTemplate(params Params) (*GetBlockTemplate, error) {
	data, err := c.send("getblocktemplate", params)
	if err != nil {
		return nil, err
	}
	var v = &GetBlockTemplate{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetLocalSolPS(params Params) (*GetLocalSolPS, error) {
	data, err := c.send("getlocalsolps", params)
	if err != nil {
		return nil, err
	}
	var v = &GetLocalSolPS{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetMiningInfo(params Params) (*GetMiningInfo, error) {
	data, err := c.send("getmininginfo", params)
	if err != nil {
		return nil, err
	}
	var v = &GetMiningInfo{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetNetworkHashPS(params Params) (*GetNetworkHashPS, error) {
	data, err := c.send("getnetworkhashps", params)
	if err != nil {
		return nil, err
	}
	var v = &GetNetworkHashPS{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetNetworkSolPS(params Params) (*GetNetworkSolPS, error) {
	data, err := c.send("getnetworksolps", params)
	if err != nil {
		return nil, err
	}
	var v = &GetNetworkSolPS{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) PrioritizeTransaction(params Params) (*PrioritizeTransaction, error) {
	data, err := c.send("prioritizetransaction", params)
	if err != nil {
		return nil, err
	}
	var v = &PrioritizeTransaction{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) SubmitBlock(params Params) (*SubmitBlock, error) {
	data, err := c.send("z_getpaymentdisclosure", params)
	if err != nil {
		return nil, err
	}
	var v = &SubmitBlock{}
	return v, json.Unmarshal(data, v)
}
