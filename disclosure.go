package komodo

import "net/http"

// ZGetPaymentDisclosure method generates a payment disclosure for a given joinsplit
// output
/*
	Params:
		txid (string, required)
		js_index (string, required)
		output_index (string, required)
		message (string, optional)
	Returns:
		http response and error
*/
func (c *Client) ZGetPaymentDisclosure(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("zgetpaymentdisclosure", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// ZValidatePaymenyDisclosure method validates the payment disclosure
/*
	Params:
		paymentdisclosure (string, required) - hex data string, with "zpd:" prefix
	Returns:
		http response and error
*/
func (c *Client) ZValidatePaymentDisclosure(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("zvalidatepaymentdisclosure", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
