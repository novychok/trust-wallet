package service

import "github.com/novychok/trustwallet"

type Parser interface {
	// last parsed block
	GetCurrentBlock() (*trustwallet.ParsedBlock, error)
	// add address to observer
	Subscribe(from, to string, tx *trustwallet.Transaction) error
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) ([]*trustwallet.Transaction, error)
	GetAddresses() map[string][]*trustwallet.Transaction
}
