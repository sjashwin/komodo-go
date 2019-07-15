package komodo

import "encoding/json"

func (c *Client) AddMultiSigAddress(params Params) (*AddMultiSigAddress, error) {
	data, err := c.send("addmultisigaddress", params)
	if err != nil {
		return nil, err
	}
	var v = &AddMultiSigAddress{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) BackupWallet(params Params) (*BackupWallet, error) {
	data, err := c.send("backupwallet", params)
	if err != nil {
		return nil, err
	}
	var v = &BackupWallet{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) DumpPrivKey(params Params) (*DumpPrivKey, error) {
	data, err := c.send("dumpprivkey", params)
	if err != nil {
		return nil, err
	}
	var v = &DumpPrivKey{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) DumpWallet(params Params) (*DumpWallet, error) {
	data, err := c.send("dumpwallet", params)
	if err != nil {
		return nil, err
	}
	var v = &DumpWallet{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) EncryptWallet(params Params) (*EncryptWallet, error) {
	data, err := c.send("encryptwallet", params)
	if err != nil {
		return nil, err
	}
	var v = &EncryptWallet{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetBalance(params Params) (*GetBalance, error) {
	data, err := c.send("getbalance", params)
	if err != nil {
		return nil, err
	}
	var v = &GetBalance{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetNewAddress(params Params) (*GetNewAddress, error) {
	data, err := c.send("GetNewAddress", params)
	if err != nil {
		return nil, err
	}
	var v = &GetNewAddress{}
	return v, json.Unmarshal(data, v)
}
