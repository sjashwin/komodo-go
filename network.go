package komodo

import "encoding/json"

func (c *Client) AddNode(params Params) error {
	_, err := c.send("addnode", params)
	return err
}

func (c *Client) ClearBanned() error {
	_, err := c.send("clearbanned", nil)
	return err
}

func (c *Client) DisconnectNode(params Params) error {
	_, err := c.send("disconnectnode", params)
	return err
}

func (c *Client) GetAddedNodeInfo(params Params) (*GetAddedNodeInfo, error) {
	data, err := c.send("getaddednodeinfo", params)
	if err != nil {
		return nil, err
	}
	var v = &GetAddedNodeInfo{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetConnectionCount() (*GetConnectionCount, error) {
	data, err := c.send("generate", nil)
	if err != nil {
		return nil, err
	}
	var v = &GetConnectionCount{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetDeprecationInfo() (*GetDeprecationInfo, error) {
	data, err := c.send("getdeprecationinfo", nil)
	if err != nil {
		return nil, err
	}
	var v = &GetDeprecationInfo{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetNetTotals() (*GetNetTotals, error) {
	data, err := c.send("getnettotals", nil)
	if err != nil {
		return nil, err
	}
	var v = &GetNetTotals{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetNetworkInfo() (*GetNetworkInfo, error) {
	data, err := c.send("getnetworkinfo", nil)
	if err != nil {
		return nil, err
	}
	var v = &GetNetworkInfo{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetPeerInfo() (*GetPeerInfo, error) {
	data, err := c.send("getpeerinfo", nil)
	if err != nil {
		return nil, err
	}
	var v = &GetPeerInfo{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) ListBanned() (*ListBanned, error) {
	data, err := c.send("listbanned", nil)
	if err != nil {
		return nil, err
	}
	var v = &ListBanned{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) Ping() error {
	_, err := c.send("ping", nil)
	return err
}

func (c *Client) SetBan(params Params) error {
	_, err := c.send("setBan", params)
	return err
}
