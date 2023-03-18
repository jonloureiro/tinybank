package usecases

import "github.com/jonloureiro/tiny-bank/app/domain/usecases/repositories"

type TinyBankUseCases struct {
	PrivateKey string

	AccountsRepo repositories.AccountsRepository
}
