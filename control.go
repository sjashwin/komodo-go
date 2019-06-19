package komodo

import "encoding/json"

func (c *Client) GetInfo() (*GetInfo, error) {
	data, err := c.send("getinfo", nil)
	if err != nil {
		return nil, err
	}
	var v = &GetInfo{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetHelp(params Params) (*GetHelp, error) {
	data, err := c.send("help", params)
	if err != nil {
		return nil, err
	}
	var v = &GetHelp{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) Stop() (*Stop, error) {
	data, err := c.send("stop", nil)
	if err != nil {
		return nil, err
	}
	var v = &Stop{}
	return v, json.Unmarshal(data, v)
}
