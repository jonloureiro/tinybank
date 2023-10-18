package inmemoryrepository

import "github.com/jonloureiro/tinybank/internal/account"

type inMemoryRepository struct {
	storageByCPF map[string]*account.Account
	storageByID  map[string]*account.Account
}

func New() *inMemoryRepository {
	return &inMemoryRepository{
		storageByCPF: make(map[string]*account.Account),
		storageByID:  make(map[string]*account.Account),
	}
}
