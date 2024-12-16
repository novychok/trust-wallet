package trustwallet

type ParsedBlock struct {
	Number string  `json:"number"`
	Result *Result `json:"result"`
}

type Result struct {
	Transactions []*Transaction `json:"transactions"`
}

type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	ChainId          string `json:"chainId"`
	From             string `json:"from"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
}
