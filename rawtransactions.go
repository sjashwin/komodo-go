package komodo

import "encoding/json"

func (c *Client) CreateRawTransactions(params Params) (*CreateRawTransactions, error) {
	data, err := c.send("createrawtransaction", nil)
	if err != nil {
		return nil, err
	}
	var v = &CreateRawTransactions{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) DecodeRawTransactions(params Params) (*DecodeRawTransactions, error) {
	data, err := c.send("decoderawtransaction", nil)
	if err != nil {
		return nil, err
	}
	var v = &DecodeRawTransactions{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) DecodeScript(params Params) (*DecodeScript, error) {
	data, err := c.send("decodescript", nil)
	if err != nil {
		return nil, err
	}
	var v = &DecodeScript{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetRawTransaction(params Params) (*GetRawTransaction, error) {
	data, err := c.send("getrawtransaction", nil)
	if err != nil {
		return nil, err
	}
	var v = &GetRawTransaction{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) SendRawTransaction(params Params) (*SendRawTransaction, error) {
	data, err := c.send("sendrawtransaction", nil)
	if err != nil {
		return nil, err
	}
	var v = &SendRawTransaction{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) SignRawTransaction(params Params) (*SignRawTransaction, error) {
	data, err := c.send("signrawtransaction", nil)
	if err != nil {
		return nil, err
	}
	var v = &SignRawTransaction{}
	return v, json.Unmarshal(data, v)
}
