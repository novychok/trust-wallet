package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/novychok/trustwallet"
	"github.com/novychok/trustwallet/repository"
)

type srv struct {
	parserRepository repository.Parser
}

var urlEndpoint = "https://ethereum-rpc.publicnode.com"

func (s *srv) GetCurrentBlock() (*trustwallet.ParsedBlock, error) {

	parsedBlockPayload := []byte(`{
		"jsonrpc": "2.0",
		"method": "eth_getBlockByNumber",
		"params": ["latest", true],
		"id": 1
	}`)

	req, err := http.NewRequest("POST", urlEndpoint, bytes.NewReader(parsedBlockPayload))
	if err != nil {
		return nil, fmt.Errorf("error to make new request: err = %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error to do the request: err = %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error to read resp body: err = %s", err)
	}

	var parsedBlock *trustwallet.ParsedBlock
	err = json.Unmarshal(body, &parsedBlock)
	if err != nil {
		return nil, fmt.Errorf("error to unmarshal parsed block: err = %s", err)
	}

	return parsedBlock, nil
}

func (s *srv) Subscribe(from, to string, tx *trustwallet.Transaction) error {

	err := s.parserRepository.Create(from, to, tx)
	if err != nil {
		return fmt.Errorf("error to subscribe the address: err = %s", err)
	}

	return nil
}

func (s *srv) GetTransactions(address string) ([]*trustwallet.Transaction, error) {

	transactions, err := s.parserRepository.GetTransactions(address)
	if err != nil {
		return nil, fmt.Errorf("error get all transactions: err = %s", err)
	}

	return transactions, nil
}

func (s *srv) GetAddresses() map[string][]*trustwallet.Transaction {
	return s.parserRepository.GetAddresses()
}

func New(parserRepository repository.Parser) Parser {
	return &srv{
		parserRepository: parserRepository,
	}
}
