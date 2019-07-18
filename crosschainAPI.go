package komodo

import "encoding/json"

func (c *Client) MigrateCreateBurnTransaction(params Params) (*MigrateCreateBurnTransaction, error) {
	data, err := c.send("migrate_createburntransaction", params)
	if err != nil {
		return nil, err
	}
	var v = &MigrateCreateBurnTransaction{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) MigrateConvertToExport(params Params) (*MigrateConvertToExport, error) {
	data, err := c.send("migrate_converttoexport", params)
	if err != nil {
		return nil, err
	}
	var v = &MigrateConvertToExport{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) MigrateCreateImportTransaction(params Params) (*MigrateCreateImportTransaction, error) {
	data, err := c.send("migrate_createimporttransaction", params)
	if err != nil {
		return nil, err
	}
	var v = &MigrateCreateImportTransaction{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) MigrateCompleteImportTransaction(params Params) (*MigrateCompleteImportTransaction, error) {
	data, err := c.send("migrate_completeimporttransaction", params)
	if err != nil {
		return nil, err
	}
	var v = &MigrateCompleteImportTransaction{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) MigrateCheckBurnTransactionSource(params Params) (*MigrateCheckBurnTransactionSource, error) {
	data, err := c.send("migrate_checkburntransaction", params)
	if err != nil {
		return nil, err
	}
	var v = &MigrateCheckBurnTransactionSource{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) MigrateCreateNotaryApprovalTransaction(params Params) (*MigrateCreateNotaryApprovalTransaction, error) {
	data, err := c.send("migrate_createnotaryapprovaltransaction", params)
	if err != nil {
		return nil, err
	}
	var v = &MigrateCreateNotaryApprovalTransaction{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) SelfImport(params Params) (*SelfImport, error) {
	data, err := c.send("selfimport", params)
	if err != nil {
		return nil, err
	}
	var v = &SelfImport{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) CalcMOM(params Params) (*CalcMOM, error) {
	data, err := c.send("calc_MoM", params)
	if err != nil {
		return nil, err
	}
	var v = &CalcMOM{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) MOMOMData(params Params) (*MOMOMData, error) {
	data, err := c.send("MoMoMdata", params)
	if err != nil {
		return nil, err
	}
	var v = &MOMOMData{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) AssetChainProof(params Params) (*AssetChainProof, error) {
	data, err := c.send("assetchainproof", params)
	if err != nil {
		return nil, err
	}
	var v = &AssetChainProof{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetNotarizationForBlock(params Params) (*GetNotarizationForBlock, error) {
	data, err := c.send("getnotarizationforblock", params)
	if err != nil {
		return nil, err
	}
	var v = &GetNotarizationForBlock{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) ScanNotarisationDB(params Params) (*ScanNotarisationDB, error) {
	data, err := c.send("scanNotarisationDB", params)
	if err != nil {
		return nil, err
	}
	var v = &ScanNotarisationDB{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetImports(params Params) (*GetImports, error) {
	data, err := c.send("getimports", params)
	if err != nil {
		return nil, err
	}
	var v = &GetImports{}
	return v, json.Unmarshal(data, v)
}

func (c *Client) GetWalletBurnTransactions(params Params) (*GetWalletBurnTransactions, error) {
	data, err := c.send("getwalletburntransactions", params)
	if err != nil {
		return nil, err
	}
	var v = &GetWalletBurnTransactions{}
	return v, json.Unmarshal(data, v)
}
