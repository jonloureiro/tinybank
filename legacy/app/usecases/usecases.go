package usecases

import "github.com/jonloureiro/tiny-bank/legacy/app/usecases/repositories"

type TinyBankUseCases struct {
	PrivateKey string

	AccountsRepo repositories.AccountsRepository
}
