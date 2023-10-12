package factories

import "github.com/jonloureiro/tiny-bank/internal/accounts/gateways/repositories"

var (
	_accountsRepositoryInMemory = repositories.NewAccountsRepositoryInMemory()
)
