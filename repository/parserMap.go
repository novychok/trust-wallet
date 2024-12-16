package repository

import (
	"sync"

	"github.com/novychok/trustwallet"
)

type mapRepository struct {
	mu         *sync.Mutex
	rw         *sync.RWMutex
	mapStorage map[string][]*trustwallet.Transaction
}

func (mr *mapRepository) Create(from, to string, tx *trustwallet.Transaction) error {

	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, ok := mr.mapStorage[from]; !ok {
		newTxStorage := make([]*trustwallet.Transaction, 0, 128)
		mr.mapStorage[from] = newTxStorage
		mr.mapStorage[from] = append(mr.mapStorage[from], tx)
	} else {
		mr.mapStorage[from] = append(mr.mapStorage[from], tx)
	}

	if _, ok := mr.mapStorage[to]; !ok {
		newTxStorage := make([]*trustwallet.Transaction, 0, 128)
		mr.mapStorage[to] = newTxStorage
		mr.mapStorage[to] = append(mr.mapStorage[to], tx)
	} else {
		mr.mapStorage[to] = append(mr.mapStorage[to], tx)
	}

	return nil
}

func (mr *mapRepository) GetTransactions(address string) ([]*trustwallet.Transaction, error) {

	mr.rw.RLock()
	defer mr.rw.RUnlock()
	transactions := mr.mapStorage[address]

	return transactions, nil
}

func (mr *mapRepository) GetAddresses() map[string][]*trustwallet.Transaction {
	return mr.mapStorage
}

func NewMapRepository() Parser {
	return &mapRepository{
		mu:         &sync.Mutex{},
		rw:         &sync.RWMutex{},
		mapStorage: make(map[string][]*trustwallet.Transaction),
	}
}
