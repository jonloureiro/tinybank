package factories

import "github.com/jonloureiro/tiny-bank/internal/accounts/gateways/presenters"

var (
	_createAccountJsonPresenter = presenters.NewCreateAccountJsonPresenter()
	_listAccountsJsonPresenter  = presenters.NewListAccountsJsonPresenter()
)
