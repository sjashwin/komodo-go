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

type GetInfo struct {
	Result struct {
		Version            int32   `json:"version"`
		ProtocolVersion    int32   `json:"protocolversion"`
		KMDVersion         int32   `json:"KMDversion"`
		Notarized          int32   `json:"notarized"`
		PerMoMHeight       int32   `json:"perMoMheight"`
		NotarizedHash      string  `json:"notarizedhash"`
		NotarizedTXID      string  `json:"notarizedtxid"`
		NotarizedTXHeight  string  `json:"notarizedtxheight"`
		KMDnotarizedHeight int32   `json:"KMDnotarizedHeight"`
		NotarizedConfirms  int32   `json:"notarizedconfirms"`
		WalletVersion      int32   `json:"walletversion"`
		Balance            float32 `json:"balance"`
		Blocks             int32   `json:"blocks"`
		LongestChain       int32   `json:"longestchain"`
		TimeOffset         int32   `json:"timeoffset"`
		TipTime            int32   `json:"tiptime"`
		Connections        int32   `json:"connections"`
		Proxy              string  `json:"proxy"`
		Difficulty         float32 `json:"difficulty"`
		TestNet            bool    `json:"testnet"`
		KeyOldestPool      int32   `json:"keyoldestpool"`
		KeyPoolSize        int32   `json:"keypoolsize"`
		RelayFee           float64 `json:"relayfee"`
		PayTXFee           float32 `json:"paytxfee"`
		Errors             string  `json:"errors"`
		Name               string  `json:"name"`
		P2PPort            int32   `json:"p2pport"`
		RPCPort            int32   `json:"rpcport"`
		Magic              int32   `json:"magic"`
		PreMine            int32   `json:"premine"`
	}
	Common
}

type GetHelp struct {
	Result string `json:"result"`
	Common
}

type Stop struct {
	Result string `json:"result"`
	Common
}

type ZGetPaymentDisclosure struct {
	Result string `json:"result"`
	Common
}

type Generate struct {
	Result []string `json:"result"`
	Common
}

type GetGenerate struct {
	Result bool `json:"result"`
	Common
}

type GetBlockSubsidy struct {
	Result struct {
		Miner float32 `json:"miner"`
	}
	Common
}

type GetBlockTemplate struct {
	Result struct {
		Version              int32         `json:"version"`
		PreviousBlockHash    string        `json:"previousblockhash"`
		FinalSaplingRootHash string        `json:"finalsaplingroothash"`
		Transactions         []Transaction `json:"transactions"`
		CoinbaseTXn          CoinbaseTXn   `json:"coinbasetxn"`
		LongPollID           string        `json:"longpollid"`
		Target               string        `json:"target"`
		MinTime              int32         `json:"mintime"`
		Mutable              []string      `json:"mutable"`
		NonceRange           string        `json:"noncerange"`
		SigOpLimit           int32         `json:"sigoplimit"`
		SizeLimit            int32         `json:"sizelimit"`
		CurTime              int32         `json:"curtime"`
		Bits                 string        `json:"bits"`
		Height               int32         `json:"height"`
	}
	Common
}

type Transaction struct {
	Data    string  `json:"data"`
	Hash    string  `json:"hash"`
	Depends []int32 `json:"depends"`
	Fee     int32   `json:"fee"`
	SigOPs  int32   `json:"sigops"`
}

type CoinbaseTXn struct {
	Data          string  `json:"data"`
	Hash          string  `json:"hash"`
	Depends       []int32 `json:"depends"`
	Fee           int32   `json:"fee"`
	SigOPs        int32   `json:"sigops"`
	CoinbaseValue int64   `json:"coinbasevalue"`
	Required      bool    `json:"required"`
}

type GetLocalSolPS struct {
	Result float64 `json:"result"`
	Common
}

type GetMiningInfo struct {
	Result struct {
		Blocks           int32   `json:"blocks"`
		CurrentBlockSize int32   `json:"currentblocksize"`
		Difficulty       float32 `json:"difficulty"`
		Errors           string  `json:"errors"`
		GenProcLimit     int32   `json:"genproclimit"`
		LocalSolPS       int32   `json:"localsolps"`
		NetworkSolPS     int32   `json:"networksolps"`
		NetworkHashPS    int32   `json:"networkhashps"`
		PooledTX         int32   `json:"pooledtx"`
		TestNet          bool    `json:"testnet"`
		Chain            string  `json:"chain"`
		Generate         bool    `json:"generate"`
		NumThreads       int32   `json:"numthreads"`
	}
	Common
}

type GetAddedNodeInfo struct {
	Result struct {
		AddedNode string    `json:"addednode"`
		Connected bool      `json:"connected"`
		Addresses []Address `json:"addresses"`
	}
	Common
}

type GetConnectionCount struct {
	Result int32 `json:"getconnectioncount"`
	Common
}

type GetDeprecationInfo struct {
	Result struct {
		Version           int32  `json:"version"`
		SubVersion        string `json:"subversion"`
		DeprecationHeight int32  `json:"deprecationheight"`
	}
	Common
}

type GetNetTotals struct {
	Result struct {
		TotalBytesRecv int32 `json:"totalbytesrecv"`
		TotalBytesSent int32 `json:"totalbytessent"`
		TimeMillis     int64 `json:"timemillis"`
	}
	Common
}

type GetNetworkInfo struct {
	Result struct {
		Version         int32     `json:"version"`
		SubVersion      string    `json:"subversion"`
		ProtocolVersion int32     `json:"protocolversion"`
		LocalServices   string    `json:"localservices"`
		TimeOffset      int32     `json:"timeoffset"`
		Connections     int32     `json:"connections"`
		Network         []Network `json:"network"`
		Relay           float64   `json:"relay"`
		// Need to find data structure.
		LocalAddresses []interface{} `json:"localaddress"`
		Warnings       string        `json:"warnings"`
	}
	Common
}

type Network struct {
	Name                       string `json:"name"`
	Limited                    bool   `json:"limited"`
	Reachable                  bool   `json:"reachable"`
	Proxy                      string `json:"proxy"`
	ProxyRandomizedCredentials string `json:"proxyrandomizedcredentials"`
}

type GetPeerInfo struct {
	Result struct {
		ID             int32   `json:"id"`
		Addr           string  `json:"addr"`
		AddrLocal      string  `json:"addrlocal"`
		Services       string  `json:"services"`
		LastSend       int32   `json:"lastsend"`
		LastRecv       int64   `json:"lastrecv"`
		BytesSent      int64   `json:"bytessent"`
		BytesRecv      int64   `json:"bytesrecv"`
		ConnTime       int64   `json:"conntime"`
		TimeOffset     int32   `json:"timeoffset"`
		PingTime       float64 `json:"pingtime"`
		Version        int32   `json:"version"`
		SubVer         string  `json:"subver"`
		InBound        bool    `json:"inbound"`
		StartingHeight int32   `json:"startingheight"`
		BanScore       int32   `json:"banscore"`
		SyncedHeaders  int32   `json:"synced_headers"`
		SyncedBlocks   int32   `json:"synced_blocks"`
		// Data type unknown
		InFlight    []interface{} `json:"inflight"`
		WhiteListed bool          `json:"whitelisted"`
	}
	Common
}

type ListBanned struct {
	Result []Banned `json:"result"`
	Common
}

type Banned struct {
	Address     string `json:"addresss"`
	BannedUntil int64  `json:"banned_until"`
}

type Address struct {
	Address   string `json:"address"`
	Connected string `json:"connected"`
}

type GetNetworkHashPS struct {
	Result int32 `json:"result"`
	Common
}

type GetNetworkSolPS struct {
	Result int32 `json:"result"`
	Common
}

type PrioritizeTransaction struct {
	Result bool `json:"result"`
	Common
}

type SubmitBlock struct {
	Result string `json:"result"`
	Common
}

type CreateRawTransactions struct {
	Result string `json:"result"`
	Common
}

type DecodeRawTransactions struct {
	Result struct {
		TXID       string        `json:"txid"`
		Size       int32         `json:"size"`
		Version    int32         `json:"version"`
		LockTime   int32         `json:"locatime"`
		Vins       []Vin         `json:"vin"`
		Vouts      []Vout        `json:"vout"`
		VJoinSplit []interface{} `json:"vjoinsplit"`
	}
	Common
}

type DecodeScript struct {
	Result struct {
		ASM  string `json:"asm"`
		Type string `json:"type"`
		P2Sh string `json:"p2sh"`
	}
	Common
}

type CreateMultiSig struct {
	Address      string `json:"address"`
	RedeemScript string `json:"redeemscript"`
	Common
}

type DecodeCCOpret struct {
	Result string  `json:"result"`
	OpRets []OpRet `json:"OpRets"`
	Common
}

type InvalidateBlock struct {
	Result string `json:"result"`
	Common
}

type ValidateAddress struct {
	IsValid      bool   `json:"isvalid"`
	Address      string `json:"address"`
	ScriptPubKey string `json:"scriptpubkey"`
	SegID        int    `json:"segid"`
	IsMine       bool   `json:"ismine"`
	IsWatchOnly  bool   `json:"iswatchonly"`
	IsScript     bool   `json:"isscript"`
	PubKey       string `json:"pubkey"`
	IsCompressed bool   `json:"iscompressed"`
	Account      string `json:"account"`
}

type VerifyMessage struct {
	bool
}

type ZValidateAddress struct {
	IsValid         bool   `json:"isvalid"`
	Address         string `json:"address"`
	PayingKey       string `json:"payingkey"`
	TransmissionKey string `json:"transmissionkey"`
	IsMine          bool   `json:"ismine"`
}

type OpRet struct {
	EvalCode string `json:"eval_code"`
	Function string `json:"function"`
}

type GetRawTransaction struct {
	Result struct {
		Hex              string        `json:"hex"`
		Txid             string        `json:"txid"`
		OverWintered     bool          `json:"overwintered"`
		Version          int32         `json:"version"`
		VersionGroupID   string        `json:"versiongroupid"`
		LockTime         int64         `json:"locktime"`
		ExpiryHeight     int64         `json:"expiryheight"`
		Vins             []Vin         `json:"vin"`
		Vouts            []Vout        `json:"vout"`
		VJoinSplit       []interface{} `json:"vjoinsplit"`
		ValueBalance     float32       `json:"valuebalance"`
		VShieldSpend     []interface{} `json:"vshieldspend"`  // Data type unknown
		VShieldOutput    []interface{} `json:"vshieldoutput"` // Data type unknown
		BlockHash        string        `json:"blockhash"`
		Height           int32         `json:"height"`
		Confirmations    int32         `json:"confirmations"`
		RawConfirmations int32         `json:"rawconfirmations"`
		Time             int64         `json:"time"`
		BlockTime        int64         `json:"blocktime"`
	}
	Common
}

type AddMultiSigAddress struct {
	Result string `json:"result"`
	Common
}

type BackupWallet struct {
	Result string `json:"result"`
	Common
}

type DumpPrivKey struct {
	Result string `json:"result"`
	Common
}

type DumpWallet struct {
	Result string `json:"result"`
	Common
}

type EncryptWallet struct {
	Result string `json:"result"`
	Common
}

type GetBalance struct {
	Result float64 `json:"result"`
	Common
}

type GetNewAddress struct {
	Result string `json:"result"`
	Common
}

type GetRawAddressChange struct {
	Result string `json:"result"`
	Common
}

type GetReceivedByAddress struct {
	Result float64 `json:"result"`
	Common
}

type GetTransaction struct {
	Result struct {
		Amount           float64   `json:"amount"`
		RawConfirmations int       `json:"rawconfirmations"`
		Confirmations    int       `json:"confirmations"`
		BlockHash        string    `json:"blockhash"`
		BlockIndex       int       `json:"blockindex"`
		BlockTime        int64     `json:"blocktime"`
		ExpiryHeight     int64     `json:"expiryheight"`
		TXID             string    `json:"txid"`
		WalletConflicts  []string  `json:"walletconflicts"`
		Time             int64     `json:"time"`
		TimeReceived     int64     `json:"timereceived"`
		VJoinSplit       []string  `json:"vjoinsplit"`
		Details          []Details `json:"details"`
		Hex              string    `json:"hex"`
	}
	Common
}

type GetUnconfirmedBalance struct {
	Result float64 `json:"result"`
	Common
}

type ImportAddress struct {
	Result string `json:"result"`
	Common
}

type ImportPrivKey struct {
	Result string `json:"result"`
	Common
}

type ImportWallet struct {
	Result string `json:"result"`
	Common
}

type KeyPoolRefill struct {
	Result string `json:"result"`
	Common
}

type ListUnlockSpent struct {
	Result []ListUnlock `json:"Result"`
	Common
}

type ListUnlock struct {
	TXID string `json:"txid"`
	VOut int    `json:"vout"`
	Common
}

type ListReceivedByAddres struct {
	Result []ReceivedAddress `json:"result"`
	Common
}

type ReceivedAddress struct {
	Address          string   `json:"address"`
	Account          string   `json:"account"`
	Amount           float64  `json:"amount"`
	RawConfirmations int      `json:"rawconfirmations"`
	TXIDs            []string `json:"txids"`
}

type ListSinceBlock struct {
	Result    []Transactions `json:"result"`
	LastBlock string         `json:"lastblock"`
	Common
}

type ListTransactions struct {
	Result    []Transactions `json:"result"`
	LastBlock string         `json:"lastblock"`
	Common
}

type ListUnspent struct {
	Result []Unspent `json:"result"`
	Common
}

type Unspent struct {
	TXID             string  `json:"txid"`
	Vout             int     `json:"vout"`
	Generated        bool    `json:"generated"`
	Address          string  `json:"address"`
	ScriptPubKey     string  `json:"scriptpubkey"`
	Amount           float64 `json:"amount"`
	Interest         float64 `json:"interest"`
	RawConfirmations int     `json:"rawconfirmations"`
	Confirmations    int     `json:"confirmations"`
	Spendable        bool    `json:"spendable"`
}

type LockUnspent struct {
	Result bool `json:"result"`
	Common
}

type ResendWalletTransactions struct {
	Result []string `json:"result"`
	Common
}

type SendMany struct {
	Result string `json:"result"`
	Common
}

type SendToAddress struct {
	Result string `json:"result"`
	Common
}

type SetPubKey struct {
	Result struct {
		Address string `json:"address"`
		IsMine  bool   `json:"ismine"`
		PubKey  string `json:"pubkey"`
	}
	Common
}

type SetTXFee struct {
	Result bool `json:"result"`
	Common
}

type SignMessage struct {
	Result string `json:"result"`
	Common
}

type ZExportKey struct {
	Result string `json:"result"`
	Common
}

type ZExportViewingKey struct {
	Result string `json:"result"`
	Common
}

type ZExportWallet struct {
	Result string `json:"result"`
	Common
}

type ZGetBalance struct {
	Result float64 `json:"result"`
	Common
}

type ZGetNewAddress struct {
	Result string `json:"result"`
	Common
}

type ZGetOperationResult struct {
	Result []OperationResult `json:"result"`
	Common
}

type ZGetOperationStatus struct {
	Result []OperationResult `json:"result"`
	Common
}

type ZGetTotalBalance struct {
	Result struct {
		Transparent string `json:"transparent"`
		Interest    string `json:"interest"`
		Private     string `json:"private"`
		Total       string `json:"total"`
	}
	Common
}

type ZImportKey struct {
	Result string `json:"result"`
	Common
}

type ZListAddress struct {
	Result []string `json:"result"`
	Common
}

type ZListOperationIDs struct {
	Result []string `json:"result"`
	Common
}

type OperationResult struct {
	ID           string   `json:"id"`
	Status       string   `json:"status"`
	CreationTime int64    `json:"creation_time"`
	Result       []string `json:"result"`
	ExectionSecs float64  `json:"execution_secs"`
	Method       string   `json:"method"`
	Params       struct {
		FromAddress string   `json:"fromaddress"`
		Amounts     []Amount `json:"amounts"`
		MineConf    int      `json:"mineconf"`
		Fee         float64  `json:"fee"`
	}
}

type Amount struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

type Transactions struct {
	Account          string   `json:"account"`
	Address          string   `json:"address"`
	Category         string   `json:"category"`
	Amount           float64  `json:"amount"`
	Vout             int      `json:"vout"`
	RawConfirmations int      `json:"rawconfirmations"`
	Confirmations    int      `json:"confirmations"`
	BlockHash        string   `json:"blockhash"`
	BlockIndex       int      `json:"blockindex"`
	BlockTime        int64    `json:"blocktime"`
	ExpiryHeight     int64    `json:"expiryheight"`
	TXID             string   `json:"txid"`
	WalletConflicts  []string `json:"walletconflicts"`
	Time             int64    `json:"time"`
	TimeReceived     int64    `json:"timereceived"`
	VJoinSplit       []string `json:"vjoinsplit"`
	Size             int      `json:"size"`
}

type Details struct {
	Account  string  `json:"account"`
	Address  string  `json:"address"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	VOut     int     `json:"vout"`
	Size     int     `json:"size"`
}

type SendRawTransaction struct {
	Result string `json:"result"`
	Common
}

type SignRawTransaction struct {
	Result struct {
		Hex       string `json:"hex"`
		Completed bool   `json:"completed"`
	}
	Common
}

type MigrateCreateBurnTransaction struct {
	Result struct {
		PayOuts   string `json:"payouts"`
		BurnTxHex string `json:"BurnTxHex"`
	}
	Common
}

type MigrateConvertToExport struct {
	Result struct {
		PayOuts  string `json:"payouts"`
		ExportTx string `json:"exportTx"`
	}
	Common
}

type MigrateCreateImportTransaction struct {
	Result struct {
		ImportTxHex string `json:"ImportTxHex"`
	}
	Common
}

type MigrateCompleteImportTransaction struct {
	Result struct {
		ImportTxHex string `json:"ImportTxHex"`
	}
	Common
}

type MigrateCheckBurnTransactionSource struct {
	Result struct {
		SourceSymbol string `json:"SourceSymbol"`
		TargetSymbol string `json:"TargetSymbol"`
		TargetCCID   string `json:"TargetCCid"`
		TxOutProof   string `json:"TxOutProof"`
		Result       string `json:"result"`
	}
	Common
}

type MigrateCreateNotaryApprovalTransaction struct {
	Result struct {
		NotaryTxHex string `json:"NotaryTxHex"`
	}
	Common
}

type SelfImport struct {
	Result struct {
		SourceTxHex string `json:"SourceTxHex"`
		ImportTxHex string `json:"ImportTxHex"`
	}
	Common
}

type CalcMOM struct {
	Result struct {
		Coin     string `json:"coin"`
		Height   int    `json:"height"`
		MOMDepth int    `json:"MoMdepth"`
		MOM      string `json:"MoM"`
	}
	Common
}

type MOMOMData struct {
	Result struct {
		Coin             string   `json:"coin"`
		KMDHeight        int64    `json:"kmdheight"`
		CCID             int      `json:"ccid"`
		MOMs             []string `json:"MoMs"`
		NotarizationHash string   `json:"notaization_hash"`
		MOMOM            string   `json:"MoMoM"`
		Data             string   `json:"data"`
	}
	Common
}

type AssetChainProof struct {
	Result string `json:"result"`
	Common
}

type GetNotarizationForBlock struct {
	Result struct {
		KMDs []KMD `json:"KMD"`
		LABS []KMD `json:"KMD"`
	}
	Common
}

type ScanNotarisationDB struct {
	Result struct {
		Height   int64  `json:"height"`
		Hash     string `json:"hash"`
		OPReturn string `json:"opreturn"`
	}
	Common
}

type GetImports struct {
	Result struct {
		Import        []Imports `json:"imports"`
		TotalImported float32   `json:"TotalImported"`
		Time          int64     `json:"time"`
	}
	Common
}

type GetWalletBurnTransactions struct {
	Result []BurnTransactions `json:"result"`
	Common
}

type Imports struct {
	TXID   string  `json:"txid"`
	Amount float64 `json:"amount"`
	Export struct {
		TXID   string  `json:"txid"`
		Amount float64 `json:"amount"`
		Source string  `json:"source"`
	}
}

type KMD struct {
	TXID      string `json:"txid"`
	Chain     string `json:"chain"`
	Height    int64  `json:"height"`
	BlockHash string `json:"blockhash"`
	Notaries  []int  `json:"notaries"`
}

type BurnTransactions struct {
	TXID         string  `json:"txid"`
	BurnedAmount float64 `json:"burnedAmount"`
	TargetSymbol string  `json:"targetSymbol"`
	TargetCCid   string  `json:"targetCCid"`
}

type Vin struct {
	TxID      string `json:"txid"`
	Vout      int32  `json:"vout"`
	ScriptSig struct {
		ASM string `json:"asm"`
		Hex string `json:"hex"`
	}
	Sequence int64 `json:"sequence"`
}

type Vout struct {
	Value    float32 `json:"value"`
	ValueSat int64   `json:"valuesat"`
	N        int32   `json:"n"`
	ScriptPubKey
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
