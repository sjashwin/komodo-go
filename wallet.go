package komodo

import "net/http"

func (c *Client) BackupWallet(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("backupwallet", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) DumpPrivKey(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("dumpprivkey", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) DumpWallet(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("dumpwallet", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) EncryptWallet(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("encryptwallet", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetAccount(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getaccount", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetAccountAddress(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getaccountaddress", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetAddressByAccount(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getaddressbyaccount", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetBalance(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getbalance", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetBalance64(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getbalance64", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetNewAddress(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getnewaddress", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetRawChangeAddress(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getrawchangeaddress", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
func (c *Client) GetReceivedByAccount(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getreceivedbyaccount", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetReceivedByAddress(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("getreceivedbyaddress", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetTransaction(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("gettransaction", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetUnconfirmedBalance() (*http.Response, error) {
	if body, err := c.generateCurl("getunconfirmedtransaction", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) GetWalletInfo() (*http.Response, error) {
	if body, err := c.generateCurl("getwalletinfo", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ImportAddress(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("importaddress", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ImportPrivKey(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("importprivkey", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ImportWallet(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("importwallet", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) KeyPoolRefill(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("keypoolrefill", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ListAccounts(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("listaccounts", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ListAddressGroupings() (*http.Response, error) {
	if body, err := c.generateCurl("listaddressgroupings", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ListLockUnspent() (*http.Response, error) {
	if body, err := c.generateCurl("listlockunspent", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ListReceivedByAddress(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("listreceivedbyaddress", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ListSinceBlock(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("listsinceblock", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ListTransaction(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("listtransaction", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ListUnspent(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("listunspent", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) LockUnspent(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("lockunspent", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ResendWalletTransaction() (*http.Response, error) {
	if body, err := c.generateCurl("resendwallettransaction", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) SendFrom(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("sendfrom", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) SendMany(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("sendmany", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) SendToAddress(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("sendtoaddress", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) SetAccount(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("setaccount", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) SetPubKey(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("setpubkey", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) SetTXFee(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("settxfee", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) SignMessage(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("signmessage", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) WalletLock(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("walletlock", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) WalletPassPhrase(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("walletpassphrase", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) WalletPassPraseChange(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("walletpassphrasechange", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZExportKey(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_exportkey", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZExportViewingKey(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_exportviewingkey", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZExportWallet(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_exportwallet", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZGetBalance(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_getbalance", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZGetNewAddress() (*http.Response, error) {
	if body, err := c.generateCurl("z_getnewaddress", nil); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZGetOperationResult(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_getoperationresult", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZGetOperationStatus(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_getoperationstatus", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZGetTotalBalance(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_gettotalbalance", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZImportKey(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_importkey", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZImportViewingKey(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_importviewingkey", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZImportWallet(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_importwallet", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZListAddresses(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_listaddresses", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZListOperationIDs(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_listoperationids", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZListReceivedByAddress(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_listreceivedbyaddress", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZListUnspent(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_listunspent", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZMergeToAddress(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_mergetoaddress", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZSendMany(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_sendmoney", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZShieldCoinbase(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("z_shieldcoinbase", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}

func (c *Client) ZCBenchmark(params Params) (*http.Response, error) {
	if body, err := c.generateCurl("zcbenchmark", params); err != nil {
		return nil, err
	} else {
		return c.send(body)
	}
}
