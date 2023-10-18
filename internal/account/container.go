package account

type Container struct {
	CreateUsecase   CreateUsecase
	CreatePresenter CreatePresenter

	ListUsecase   ListUsecase
	ListPresenter ListPresenter

	GetBalanceUsecase   GetBalanceUsecase
	GetBalancePresenter GetBalancePresenter
}
