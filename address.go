package komodo

import "encoding/json"

// GetAddressBalance method returns the confirmed balance for an address, or addresses.
// It requires addressindex to be enabled.
/*
	Params:
		address (string) - the address
	Returns:
		GetWalletInfo
		error
*/
func (c *Client) GetAddressBalance(params Params) (*GetWalletInfo, error) {
	data, err := c.send("getaddressbalance", params)
	if err != nil {
		return nil, err
	}
	var v = &GetWalletInfo{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetAddressDeltas(params Params) (*GetAddressDeltas, error) {
	data, err := c.send("getaddressdeltas", params)
	if err != nil {
		return nil, err
	}
	var v = &GetAddressDeltas{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetAddressMempool(params Params) (*GetAddressMemPool, error) {
	data, err := c.send("getaddressmempool", params)
	if err != nil {
		return nil, err
	}
	var v = &GetAddressMemPool{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetAddressTXIDS(params Params) (*GetAddressTXIDs, error) {
	data, err := c.send("getaddresstxids", params)
	if err != nil {
		return nil, err
	}
	var v = &GetAddressTXIDs{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetAddressUTXOS(params Params) (*GetAddressUTXOs, error) {
	data, err := c.send("getaddressutxos", params)
	if err != nil {
		return nil, err
	}
	var v = &GetAddressUTXOs{}
	return v, json.Unmarshal(data, v)
}

// TODO: needs to be tested. No address is passed
func (c *Client) GetSnapShot() (*GetSnapshot, error) {
	data, err := c.send("getsnapshot", nil)
	if err != nil {
		return nil, err
	}
	var v = &GetSnapshot{}
	return v, json.Unmarshal(data, v)
}
