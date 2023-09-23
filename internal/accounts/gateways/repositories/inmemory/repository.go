package inmemory

import "github.com/jonloureiro/tiny-bank/internal/accounts"

type accountsRepositoryInMemory struct {
	storageByCPF map[string]*accounts.Account
	storageByID  map[string]*accounts.Account
}

func NewAccountsRepositoryInMemory() *accountsRepositoryInMemory {
	return &accountsRepositoryInMemory{
		storageByCPF: make(map[string]*accounts.Account),
		storageByID:  make(map[string]*accounts.Account),
	}
}
