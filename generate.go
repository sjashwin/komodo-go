package komodo

import "net/http"

// GetGenerate method returns a boolean value indicating the server's mining status
// Returns http response and error
func (c *Client) GetGenerate() (*http.Response, error) {
	if body, err := c.generateCurl("getgenerate", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// SetGenerate method allows the user to set the generate property in the coin daemon to true
// or false, this turning generation (mining/staking) on or off.
// Generation is limited to genproclimit processors. Set genproclimit to -1 to use maximum available
// processors
/*
	Params:
		generate (boolean, required) - set to true to turn on generation; set to off to turn off generation
		genproclimit (numeric, optional) - set the processor limit for when generation is on; use value
		"-1" for unlimited
*/
func (c *Client) SetGenerate(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("setgenerate", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
