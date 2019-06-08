package komodo

import "encoding/json"

func (c *Client) Generate(params Params) (*Generate, error) {
	data, err := c.send("generate", params)
	if err != nil {
		return nil, err
	}
	var v = &Generate{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetGenerate(params Params) (*GetGenerate, error) {
	data, err := c.send("getgenerate", params)
	if err != nil {
		return nil, err
	}
	var v = &GetGenerate{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) SetGenerate(params Params) error {
	_, err := c.send("setgenerate", params)
	return err
}
