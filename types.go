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

type CoinSupply struct {
	Result struct {
		Result string  `json:"result"`
		Coin   string  `json:"coin"`
		Height int32   `json:"height"`
		Supply float32 `json:"supply"`
		ZFunds float32 `json:"zfunds"`
		Sprout float32 `json:"sprout"`
		Total  float32 `json:"total"`
	}
	Common
}

type GetBestBlockHash struct {
	Result string `json:"result"`
	Common
}

type GetBlock struct {
	Result struct {
		Hash              string   `json:"hash"`
		Confirmations     int32    `json:"confirmations"`
		RawConfirmations  int32    `json:"rawconformations"`
		Size              int32    `json:"size"`
		Height            int32    `json:"height"`
		Version           int32    `json:"version"`
		Merkleroot        string   `json:"merkleroot"`
		Tx                []string `json:"tx"`
		Time              int32    `json:"time"`
		Nonce             int32    `json:"none"`
		Bits              string   `json:"bits"`
		Difficulty        int32    `json:"difficulty"`
		PreviousBlockHash string   `json:"previousblockhash"`
		NextBlockHash     string   `json:"nextblockhash"`
	}
	Hex string `json:"hex"`
	Common
}

type GetBlockchainInfo struct {
	Result struct {
		Chain                string     `json:"chain"`
		Blocks               int32      `json:"blocks"`
		Headers              int32      `json:"headers"`
		BestBlockHash        string     `json:"bestblockhash"`
		Difficulty           int32      `json:"difficulty"`
		VerificationProgress int32      `json:"verificationprogress"`
		ChainWork            string     `json:"chainwork"`
		Pruned               bool       `json:"pruned"`
		SizeOnDisk           int32      `json:"sizeondisk"`
		Commitments          int32      `json:"commitments"`
		SoftFork             []SoftFork `json:"softfork"`
	}
	Common
}

type GetBlockCount struct {
	Result int32 `json:"result"`
	Common
}

type GetBlockHash struct {
	Result string `json:"result"`
	Common
}

type GetBlockHashes struct {
	Result struct {
		Hash      string `json:"hash"`
		BlockHash string `json:"blockhash"`
		LogicAlts int32  `json:"logicalts"`
	}
	Common
}

type GetBlockHeader struct {
	Result struct {
		Hash              string `json:"hash"`
		Confirmations     int32  `json:"confirmations"`
		RawConfirmations  int32  `json:"rawconfirmations"`
		Height            int32  `json:"height"`
		Version           int32  `json:"version"`
		Merkleroot        string `json:"merkleroot"`
		Time              int32  `json:"time"`
		Nonce             int32  `json:"nonce"`
		Bits              string `json:"bits"`
		Difficulty        int32  `json:"difficulty"`
		PreviousBlockHash string `json:"previousblockhash"`
		NextBlockHash     string `json:"nextblockhash"`
	}
	Data string `json:"data"`
	Common
}

type GetChainTips struct {
	// Needs to be reviewd carefully
	Result struct {
		Height    int32  `json:"height"`
		Hash      string `json:"hash"`
		BranchLen int32  `json:"branchlen"`
		Status    string `json:"status"`
	}
	Common
}

type GetChainTXStats struct {
	Result struct {
		Time                 int32  `json:"time"`
		TXCount              int32  `json:"txcount"`
		WindowFinalBlockHash string `json:"window_final_block_hash"`
		WindowBlockCount     int32  `json:"window_block_count"`
		WindowTXCount        int32  `json:"window_tx_count"`
		WindowInterval       int32  `json:"window_interval"`
		TXRate               int32  `json:"txrate"`
	}
	Common
}

type GetDifficulty struct {
	Result float64 `json:"result"`
	Common
}

type GetLastSegIDStakes struct {
	Result struct {
		NotSet  int32    `json:"notset"`
		POW     int32    `json:"PoW"`
		PoSPerc int32    `json:"PoSPerc"`
		SegIDs  struct{} `json:"segIDs"`
		N       int32    `json:"n"`
	}
	Common
}

type GetMemPoolInfo struct {
	Result struct {
		Size  int32 `json:"size"`
		Bytes int32 `json:"bytes"`
		Usage int32 `json:"usage"`
	}
	Common
}

type GetRawMemPool struct {
	Result struct {
		Transaction []TransactionIDs
	}
	TransactionID string `json:"transation_id"`
	Common
}

type TransactionIDs struct {
	Size             int32    `json:"size"`
	Fee              float32  `json:"feee"`
	Time             int32    `json:"time"`
	Height           int32    `json:"height"`
	StartingPriority float64  `json:"startingpriority"`
	CurrentPriority  float64  `json:"currentpriority"`
	Depends          []string `json:"depends"`
}

type GetSpentInfo struct {
	Result struct {
		TxID   string `json:"txid"`
		Index  int32  `json:"index"`
		Height int32  `json:"height"`
	}
	Common
}

type GetTXOut struct {
	Result struct {
		BestBlock        string  `json:"bestblock"`
		Confirmations    int32   `json:"confirmations"`
		RawConfirmations int32   `json:"rawconfirmations"`
		Value            float32 `json:"value"`
		ScriptPubKey
		Version  int32 `json:"version"`
		Coinbase bool  `json:"coinbase"`
	}
	Common
}

type GetTXOutProof struct {
	Result string `json:"data"`
	Common
}

type GetTXOutSetInfo struct {
	Result struct {
		Height          int32   `json:"height"`
		BestBlock       string  `json:"bestblock"`
		Tranasctions    int32   `json:"transactions"`
		TXOuts          int32   `json:"txouts"`
		BytesSerialized int32   `json:"bytes_serialized"`
		HashSerialized  string  `json:"hash_serialized"`
		TotalAmount     float32 `json:"total_amount"`
	}
	Common
}

type KVSearch struct {
	Result struct {
		Coin          string `json:"coin"`
		CurrentHeight int32  `json:"currentheight"`
		Key           string `json:"key"`
		KeyLen        int32  `json:"keylen"`
		Owner         string `json:"owner"`
		Height        int32  `json:"height"`
		Expiration    int32  `json:"expiration"`
		Flags         int32  `json:"flags"`
		Value         string `json:"value"`
		ValueSize     int32  `json:"valuesize"`
	}
	Common
}

type KVUpdate struct {
	Result struct {
		Coin          string  `json:"coin"`
		CurrentHeight int32   `json:"currentheight"`
		Key           string  `json:"key"`
		KeyLen        int32   `json:"keylen"`
		Owner         string  `json:"owner"`
		Height        int32   `json:"height"`
		Expiration    int32   `json:"expiration"`
		Flags         int32   `json:"flags"`
		Value         string  `json:"value"`
		ValueSize     int32   `json:"valuesize"`
		Fee           float32 `json:"fee"`
		TXID          string  `json:"txid"`
	}
	Common
}

type MinerIDs struct {
	Result struct {
		Mined       []Mined `json:"mined"`
		NumNotaries int32   `json:"numnotaries"`
	}
}

type Notaryies struct {
	Result struct {
		Notaries    []Notary `json:"notaries"`
		NumNotaries int32    `json:"numnotaries"`
		Height      int32    `json:"height"`
		TimeStamp   int32    `json:"timestamp"`
	}
	Common
}

type VerifyChain struct {
	Result bool `json:"result"`
	Common
}

type VerifyTXOutProof struct {
	Result []string `json:"result"`
	Common
}

type Mined struct {
	NotaryID   int32  `json:"notaryid"`
	KMDAddress string `json:"KMDaddress"`
	PubKey     string `json:"pubkey"`
	Blocks     int32  `json:"blocks"`
}

type Notary struct {
	PubKey     string `json:"pubkey"`
	BTCAddress string `json:"BTCaddress"`
	KMDAddress string `json:"KMDaddress"`
}

type ScriptPubKey struct {
	ASM       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int32    `json:"reqsigs"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses"`
}

type SoftFork struct {
	ID      string `json:"id"`
	Version int32  `json:"version"`
	Enforce struct {
		Status   bool  `json:"status"`
		Found    int32 `json:"found"`
		Required int32 `json:"required"`
		Window   int32 `json:"window"`
	}
	Reject struct {
		Status   bool  `json:"status"`
		Found    int32 `json:"found"`
		Required int32 `json:"required"`
		Window   int32 `json:"window"`
	}
	// To Do : upgrades - could not find the data structure or type.
	Consensus struct {
		ChainTip  string `json:"chaintip"`
		NextBlock string `json:"nextblock"`
	}
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
