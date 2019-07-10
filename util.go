package komodo

import "encoding/json"

func (c *Client) CreateMultiSig(params Params) (*CreateMultiSig, error) {
	data, err := c.send("createmultisig", params)
	if err != nil {
		return nil, err
	}
	var v = &CreateMultiSig{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) DecodeCCOpret(params Params) (*DecodeCCOpret, error) {
	data, err := c.send("decodeccopret", params)
	if err != nil {
		return nil, err
	}
	var v = &DecodeCCOpret{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) InvalidateBlock(params Params) (*InvalidateBlock, error) {
	data, err := c.send("invalidateblock", params)
	if err != nil {
		return nil, err
	}
	var v = &InvalidateBlock{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) ReconsiderBlock(params Params) error {
	_, err := c.send("reconsiderblock", params)
	return err
}

func (c *Client) ValidateAddress(params Params) (*ValidateAddress, error) {
	data, err := c.send("validateaddress", params)
	if err != nil {
		return nil, err
	}
	var v = &ValidateAddress{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) VerifyMessage(params Params) (*VerifyMessage, error) {
	data, err := c.send("verifymessage", params)
	if err != nil {
		return nil, err
	}
	var v = &VerifyMessage{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) ZValidateAddress(params Params) (*ZValidateAddress, error) {
	data, err := c.send("z_validateaddress", params)
	if err != nil {
		return nil, err
	}
	var v = &ZValidateAddress{}
	return v, json.Unmarshal(data, v)
}
