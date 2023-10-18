package factory

import (
	"github.com/jonloureiro/tinybank/internal/account"
	"github.com/jonloureiro/tinybank/internal/account/app/usecase"
	"github.com/jonloureiro/tinybank/internal/account/gateway/inmemoryrepository"
	"github.com/jonloureiro/tinybank/internal/account/gateway/jsonpresenter"
)

var singletonContainer *account.Container

// Container returns the singleton account container
func Container() account.Container {
	if singletonContainer != nil {
		return *singletonContainer
	}

	var (
		repository = inmemoryrepository.New()
	)

	var (
		createUsecase     = usecase.NewCreate(repository)
		listUsecase       = usecase.NewList(repository)
		getBalanceUsecase = usecase.NewGetBalance(repository)
	)

	var (
		createPresenter     = jsonpresenter.NewCreate()
		listPresenter       = jsonpresenter.NewList()
		getBalancePresenter = jsonpresenter.NewGetBalance()
	)

	singletonContainer = &account.Container{
		CreateUsecase:   createUsecase,
		CreatePresenter: createPresenter,

		ListUsecase:   listUsecase,
		ListPresenter: listPresenter,

		GetBalanceUsecase:   getBalanceUsecase,
		GetBalancePresenter: getBalancePresenter,
	}

	return *singletonContainer
}
