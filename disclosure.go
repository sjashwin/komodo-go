package komodo

import "net/http"

func (c *Client) ZGetPaymentDisclosure(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("zgetpaymentdisclosure", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZValidatePaymentDisclosure(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("zvalidatepaymentdisclosure", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
