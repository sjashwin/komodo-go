package komodo

import "net/http"

// AddNode method attempts to add or remove a node from the addnode lsit, or to make a
// single attempt to connect to a node.
/*
	Params:
		node (string, required) - the node (see getperrinfo for nodes)
		command (string, required) - 'add' to add a node to the list, 'onetry' to try a connection
		to the node once
	Returns:
		None
*/
func (c *Client) AddNode(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("addnode", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// ClearBanned method clears all banned IPs.
// Returns None
func (c *Client) ClearBanned() (*http.Response, error) {
	if body, err := c.generateCurl("clearbanned", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// DisconnectNode method instructs the daemon to immediately disconnect from the specific node.
// Use getpeerinfo to determine the result
/*
	Params:
		node (string, required) - the node's address (see getpeerinfo for nodes)
	Returns:
		None
*/
func (c *Client) DisconnectNode(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("disconnectnode", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetAddedNodeInfo method returns information about the given added node, or all added nodes.
// If dns is set to false, only a list of added nodes is returned. Otherwise, connection information is also
// provided.
/*
	Params:
		dns (boolean, required) - if false, only a list of added nodes woll be provided; otherwise, connection information
		is also provided.
		node (string, optional) - if provided, the method returns information about this specific node;
		otherwise, all nodes are returned.
	Returns:
		http response and error
*/
func (c *Client) GetAddedNodeInfo(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getaddednodeinfo", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetConnectionCount method returns the number of connection to other nodes.
// Returns http response and error
func (c *Client) GetConnectionCount(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getconnectioncount", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetDeprecationInfo method returns an object containing currenct version and deprecation block height
// This method is only available to the KMD mainnet.
// Returns http response and error.
func (c *Client) GetDeprecationInfo() (*http.Response, error) {
	if body, err := c.generateCurl("getdeprecationinfo", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetNetTotals method returns information about network traffic, includin bytes in, bytes out and current time.
// Returns http response and error
func (c *Client) GetNetTotals() (*http.Response, error) {
	if body, err := c.generateCurl("getnettotals", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetNetowrkInfo method returns an object containing various states info regardin p2p
// networking.
// Returns http response and error
func (c *Client) GetNetworkInfo() (*http.Response, error) {
	if body, err := c.generateCurl("getnetworkinfo", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// GetPeerInfo method returns data about each connected network nodes as a json array of objects.
// Returns http response and error
func (c *Client) GetPeerInfo() (*http.Response, error) {
	if body, err := c.generateCurl("getpeerinfo", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// ListBanned method lists all banned IP addresses and subnets
// Returns http response and error
func (c *Client) ListBanned() (*http.Response, error) {
	if body, err := c.generateCurl("listbanned", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// Ping method requests that a ping be sent to all other nodes, to measure pind times.
// Results provided in getpeerinfo, pingtime and pingwait fields are decimal seconds.
// The ping command is handled in queue with all other commands, so it measures processing
// backlogs, not just network ping.
// Returns http response and error
func (c *Client) Ping() (*http.Response, error) {
	if body, err := c.generateCurl("ping", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

// SetBan method attemps to add or remove an IP address (and subnet, if indicated) from the banned list
/*
	Params:
		ip(/netmask) (string, ip required) - the IP/subnet (see getpeerinfo for node ip) with an optional
		netmask (default is /32 = single ip)
		command (string, required) - use "add" to add an IP/subnet to the list, or "remove" to remove an
		IP/subnet from the list
		bantime (numberic, optional) - indicates how long the ip is banned (or until when, if [absolute] is
		set). 0 or empty measn the ban is using the default time of 24h, which can also be overwritten using the
		-bantime runtime parameter
		absolute (boolean, optional) - if set to true, the bantime mest be an absolute timestamp (in seconds)
		since epoch(Jan 1970 GMT)
	Returns:
			http response and error
*/
func (c *Client) SetBan(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("setban", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
