package komodo

import "encoding/json"

func (c *Client) GetInfo(params Params) (*GetInfo, error) {
	data, err := c.send("getinfo", params)
	if err != nil {
		return nil, err
	}
	var v = &GetInfo{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetHelp(params Params) (*GetHelp, error) {
	data, err := c.send("gethelp", params)
	if err != nil {
		return nil, err
	}
	var v = &GetHelp{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) Stop(params Params) (*Stop, error) {
	data, err := c.send("gethelp", params)
	if err != nil {
		return nil, err
	}
	var v = &Stop{}
	return v, json.Unmarshal(data, v)
}
