package komodo

import "net/http"

// CoinSupply method returns the coin supply information for the indicated block
// height. If no height is given, the method defaults to the blockchain's current height.
/*
	Params:
		height (integer, optional) - the desired block height
	Returns:
		http response and error
*/
func (c *Client) CoinSupply(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("coinsupply", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetBlockHash method returns the hash of the best (tip) block in the longest block chain.
// Returns http response and error.
func (c *Client) GetBlockHash() (*http.Response, error) {
	if body, err := c.generateCurl("getblockhash", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetBlock method returns the block relevant state information. The verbose input is optional.
// The default value is true, and it will return a json object with information about the indicated
// block. If verbose is false, the command returns a string that is serialized hex-encoded data for
// the indicated block.
/*
	Params:
		hash or height (string or number, respectively) - the block hash or the block height
		verbose (boolean, optional, default=true) - true returns a json object, flase returns hex
		encoded
	Returns:
		http response and error
*/
func (c *Client) GetBlock(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getblock", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetBlockchainInfo method returns a json object containing the state information about blockchain
// processing
// Returns http response and error
func (c *Client) GetBlockchainInfo(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getblockchaininfo", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetBlockCount returns number of blocks in the best valid block chain.
// returns http response and error
func (c *Client) GetBlockCount() (*http.Response, error) {
	if body, err := c.generateCurl("getblockcount", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetBlockHash method returns the hash of the indicated block index, according to the best
// blockchain at the time provided.
/*
	Params:
		index (numberic, required) - the block index
	Returns:
		http response and error
*/
func (c *Client) GetBlockHashes(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getblockhashes", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetBlockHeader method returns information about the indicated block. The verbose input is optional.
// If verbose is false,  the method returns a string that is serialized, hex-encoded data for the indicated
// blockheader. If verbose is true, the method returns a json object with information about the indicated header.
/*
	Params:
		hash (string, required) - the block hash
		verbose (boolean, optional, default=true) - true returns a json object, false returns hex-encoded data
	Returns:
		http resposne and error
*/
func (c *Client) GetBlockHeader(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getblockheader", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetChainTips method returns informaito about all known tips in the block tree,
// including the main chain and any orphaned branches.
// Returns http response and error
func (c *Client) GetChainTips() (*http.Response, error) {
	if body, err := c.generateCurl("getchaintips", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetChainTXStats returns statistics about the total number and rate of transactions in the chain.
/*
	Params:
		nblocks (number, optional) - the number of blocks in the averaging window
		blockhash (string, optional) - the hash of the block which ends the window
	Returns:
		http reponse and error
*/
func (c *Client) GetChainTXStats() (*http.Response, error) {
	if body, err := c.generateCurl("getchaintxstats", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetDifficulty method returns the proof-of-work difficulty as a multiple of the minimum difficulty.
// Returns http response and error
func (c *Client) GetDifficulty(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getdifficulty", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetLastSeGgIDStakes method returns an object containing the number of blocks staked by each
// segid in the last X number of blocks, where the value of X is equal to the indicated depth.
/*
	Params:
		depth (number, required) - the number of blocks to scan, starting from the current
		height and working backwards.
	Returns:
		http response and error
*/
func (c *Client) GetLastSegIDStakes(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getlastsegidstakes", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetMemPoolInfo method returns details on the active state of the transaction memory pool.
// Returns http response and error
func (c *Client) GetMemPoolinfo(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getmempoolinfo", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetRawMemPoolInfo method returns all transaction ids in the memory pool as a json array
// of transaction ids. The verbose input is optional and is false by default. When it is true,
// the method instead returns a json object with various related data.
/*
	Params:
		verbose (boolean, optional, default=false) - true for a json object, false for a json array
		of transaction ids.
	Returns:
		http response and error
*/
func (c *Client) GetRawMemPool(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getrawmempool", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GeSpentInfo method returns the transaction id and index where the given output is spent.
// the method requires spentindex to be enabled.
/*
	Params:
		txid (string) - the transaction id
		index (number) - the spending input index
*/
func (c *Client) GetSpentInfo(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getspentinfo", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetTXOut method returns details about the unspent transaction output.
/*
	Params:
		txid (string, required) - the transaction id
		vout (numberic, required) - the vout value
		includedmempool (boolean, optional) - whether to include the mempool
	Returns:
		http response and error
*/
func (c *Client) GetTXOut(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("gettxout", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetTXOutProof method returns the hex-encoded proof showing that the indicated transaction
// was included in a block.
/*
	Params:
		txid (string) - a transaction hash
		blockhash (string, optional) - if specified, the method looks for the relevant transaction
		id in this block hash
	Returns:
		http response and error
*/
func (c *Client) GetTXOutProof(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("gettxoutproof", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetTXOutSetInfo method returns statistics about the unspent transaction output set.
// Returns http response and error
func (c *Client) GetTXOutSetInfo() (*http.Response, error) {
	if body, err := c.generateCurl("gettxoutsetinfo", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// KVSearch method searches for a key stored via the kvupdate command
/*
	Params:
		key (string, required) - the key for which the user desires to search the chain
	Returns:
		http response and error
*/
func (c *Client) KVSearch(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("kvsearch", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// KVUpdate method stores a key/value pair via OP_RETURN. This deature is available only for asset chains.
// The maximum value memory size is 8KB.
/*
	Params:
		key (string, required) - key (should be unique)
		value (string, required) - value
		days (numberic, required) - amount of days before the key expires (1440 blocks/day)l minimum 1 day
		passphrase (string, optional) - passphrase required to update this key
	Returns:
		http response and error
*/
func (c *Client) KVUpdate(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("kvupdate", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// MinerIDs method returns information about the notary nodes and external nminers at a specific block height.
// The response will calculate results according to the 2000 blocks proceeding the indicated "height" block.
/*
	Params:
		height (number) - the block height for the query
	Returns:
		http response and error
*/
func (c *Client) MinerIDs(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("minerids", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// Notaries method returns the public key, BTC address and KMD address for each Komodo notary node.
// Either or both of the height and timestamp parameters will suffice.
/*
	Params:
		height (number) - the block height desired for the query
		timestamp (number) - the timestamp of th block desired for the query
	Returns:
		http response and error
*/
func (c *Client) Notaries(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("notaires", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// VerifyChain method verifies the coin deamon's blockchain database.
/*
	Params:
		checklevel (numberic, optional, 0-4, default=3) - indicated the thoroughness of block verification
		numblocks (numberic, optional, default=288, O=all) - indicated the number of blocks to verify
	Returns:
		http response and error
*/
func (c *Client) VerifyChain(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("verifychain", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// VerifyTXOutProof method verifies that a proof points to a transaction in a block. It returns
// the transaction to which the proof is committed, or it will throw an rpc error if the block is not in the
// current best chain
/*
	Params:
		proof_string (string, required) - the hex-encoded proof generated by gettexoutproof
	Returns:
		http response and error
*/
func (c *Client) VerifyTXOutProof(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("verifytxoutproof", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
