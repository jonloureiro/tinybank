package factories

import "github.com/jonloureiro/tiny-bank/internal/accounts/gateways/frameworks"

func AccountsFactory() frameworks.Routes {
	return frameworks.Routes{
		CreateAccountUsecase:   _createAccountUsecase,
		CreateAccountPresenter: _createAccountJsonPresenter,

		ListAccountsUsecase:   _listAccountsUsecase,
		ListAccountsPresenter: _listAccountsJsonPresenter,

		GetAccountBalanceUsecase:   _getAccountBalance,
		GetAccountBalancePresenter: _getAccountBalanceJson,
	}
}
