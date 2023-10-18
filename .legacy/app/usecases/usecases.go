package usecases

import "github.com/jonloureiro/tinybank/legacy/app/usecases/repositories"

type TinyBankUseCases struct {
	PrivateKey string

	AccountsRepo repositories.AccountsRepository
}
