package komodo

type GetWalletInfo struct {
	Result struct {
		WalletVersion      int32   `json:"walletversion"`
		Balance            float32 `json:"balance"`
		UnconfirmedBalance float32 `json:"unconfirmed_balance"`
		ImmatureBalance    float32 `json:"immature_balance"`
		TXCount            int32   `json:"txcount"`
		KeyPoolOldest      int64   `json:"keypoololdest"`
		KeyPoolSize        int32   `json:"keypoolsize"`
		PatTXFee           float32 `json:"paytxfee"`
		SeedFP             string  `json:"seedfp"`
	}
	Common
}

type GetAddressBalance struct {
	Result struct {
		Balance  int32 `json:"balance"`
		Received int32 `json:"received"`
	}
	Common
}

type GetAddressDeltas struct {
	Result struct {
		Satoshis int32  `json:"satoshis"`
		TXID     string `json:"txid"`
		Index    int32  `json:"index"`
		Height   int32  `json:"height"`
		Address  string `json:"address"`
	}
	Common
}

type GetAddressMemPool struct {
	Result struct {
		Address   string `json:"address"`
		TXID      string `json:"txid"`
		Index     int32  `json:"index"`
		Satoshis  int32  `json:"satoshis"`
		Timestamp int32  `json:"timestamp"`
		PrevTXID  string `json:"prevtxid"`
		PrevOut   string `json:"prevout"`
	}
	Common
}

type GetAddressTXIDs struct {
	Result struct {
		TransactionID string `json:"transaction_id"`
	}
	Common
}

type GetAddressUTXOs struct {
	Result struct {
		Address     string `json:"address"`
		TXID        string `json:"txid"`
		Height      int32  `json:"height"`
		OutputIndex int32  `json:"outputIndex"`
		Script      string `json:"script"`
		Satoshis    int32  `json:"satoshis"`
	}
	Common
}

type GetSnapshot struct {
	Result struct {
		Addresses      Addresses `json:"addresses"`
		Address        string    `json:"addr"`
		Amount         float32   `json:"amount"`
		Total          float32   `json:"total"`
		Average        int32     `json:"average"`
		UTXOS          int32     `json:"utxos"`
		TotalAddresses int32     `json:"total_addresses"`
		StartHeight    int32     `json:"start_height"`
		EndingHeight   int32     `json:"ending_height"`
		StartTime      int32     `json:"start_time"`
		EndTime        int32     `json:"end_time"`
	}
	Common
}

type Addresses struct {
	Address string `json:"addr"`
	Amount  string `json:"amount"`
	SegID   int32  `json:"segid"`
}
type Common struct {
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}
