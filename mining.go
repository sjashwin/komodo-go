package komodo

import "net/http"

// GetBlockSubsidy method returns the block-subsidy reward. The resulting calculation takes into
// account the mining slow start. This method can be used in conjunction with custom mining rewards
// designed by the developers of a KMD-base asset chain.
/*
	Params:
		height (numeric, optional) - the block height; if the block height is not
		provided, the method defaults to the current height of the chain
	Returns:
		http response and error
*/
func (c *Client) GetBlockSubsidy(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getblocksubsidy", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetBlockTemplate nethod returns a data that is necessary to construct a block.
// If the request parameter includes a mode key, it is used to explicitly sleect between the default
// 'templates' request, a 'proposal' or 'disablecb'.
// A Note on Unique Mining Curcumstances
// There are many features on the komodo ecosystem that can make an asset chain's daemon produce a
// non-standard coinbase transactions. Examples include an asset chain parameter that creates new coins
// for a specific pubkey in every block or a CC module that adds outputs to the coinbase transaction.
// This can be dealt using a mode called `disablecb`
/*
	Params:
		jsonrequestobject:{...} (string, optional) - a json object in the following spec
		mode (string, optional) - this must be set to "template" or omitted
		capabilities (array, optional) - a list of strings
		support (string) - client side supported features: "longpoll", "coinbasetxn", "coinbasevalue",
		"proposal", "serverlist", "workid"
	Returns
		http response and error
*/
func (c *Client) GetBlockTemplate(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getblocktemplate", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetLocalSolPS method returns the average local solutions per second since this node was started
// Returns http response and error
func (c *Client) GetLocalSolPS() (*http.Response, error) {
	if body, err := c.generateCurl("getlocalsolps", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetMiningInfo method returns a json object containing mining-related information
// Returns http response and error
func (c *Client) GetMiningInfo() (*http.Response, error) {
	if body, err := c.generateCurl("verifytxoutproof", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetNetworkHashPS has been deprecated use `GetNetworkSolPS` instead
func (c *Client) GetNetworkHashPS(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getnetworkhashps", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetNetworkSolPS method returns the estimated network solutions per second based on the last n blocks
// Pass in blocks to override the default number of blocks. Use -1 to calculate according to the
// relevant difficulty averaging window. Pass in height to estimate the network speed at the time when a
// certain block was found.
/*
	Params:
		blocks (number, optional, default=20) - the number of blocks, use -1 to calculate accoring to the relevant
		difficulty averaging window.
		height (numberic. optional, default=-1) - the block height that corresponds to the requested data
	Returns:
		http response and error
*/
func (c *Client) GetNetworkSolPS(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getnetworksolps", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// PrioritiseTransaction method instructs the daemon to accept the indicated transaction into mined blocks at
// a higher (or lower) priority. The transaction selection algorithm considers the transaction as it would have
// a higher priority.
/*
	Params:
		transaction_id (string, required) - the transaction id
		priority_delta (numeric, required) - the priority to add or subtract (if negative). The
		transaction selection algorithm assigns the tx a higher or lower priority. The transaction priority
		calculation: coinage * value_in_satoshi / txsize
		fee_delta (numberic, required) - the fee value in stoshis to add or subtract (if negative); the fee
		is not actually paid, only the algorithm for selecting transaction into a block considers the
		transaction as if it paid a higher (or lower) fee.
	Returns:
		http response and error
*/
func (c *Client) PrioritiseTransaction(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("prioritisetransaction", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// SubmitBlock method instructs the daemon to propose a new block to the network.
// The jsonparmeters object parameter is currently ignored.
/*
	Params:
		hexdata (string, required) - the hex-encoded block data to submit
		jsonparamtersobject:{...} (string, optional) - object of optional parameters
		workid (string, sometimes optional) - if the server provides a workid, it MUST
		be included with submission
*/
func (c *Client) SubmitBlock(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("submitblock", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
