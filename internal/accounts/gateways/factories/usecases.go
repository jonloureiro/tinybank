package factories

import "github.com/jonloureiro/tiny-bank/internal/accounts/app"

var (
	_createAccountUsecase = app.NewCreateAccountUsecase(_accountsRepositoryInMemory)
	_listAccountsUsecase  = app.NewListAccountsUsecase(_accountsRepositoryInMemory)
)
