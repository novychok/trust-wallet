package repository

import "github.com/novychok/trustwallet"

type Parser interface {
	Create(from, to string, tx *trustwallet.Transaction) error
	GetTransactions(address string) ([]*trustwallet.Transaction, error)
	GetAddresses() map[string][]*trustwallet.Transaction
}
