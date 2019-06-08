package komodo

import "encoding/json"

func (c *Client) ZGetPaymentDisclosure(params Params) (*ZGetPaymentDisclosure, error) {
	data, err := c.send("z_getpaymentdisclosure", params)
	if err != nil {
		return nil, err
	}
	var v = &ZGetPaymentDisclosure{}
	return v, json.Unmarshal(data, v)
}
