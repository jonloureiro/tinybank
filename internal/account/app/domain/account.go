package domain

import (
	"time"

	"github.com/jonloureiro/tinybank/internal/account"
	"github.com/jonloureiro/tinybank/internal/account/app/vo"
)

const (
	_nameMinLength  = 3
	_secretMinSize  = 6
	_initialBalance = 100_00
)

type _account struct {
	id        vo.AccountID
	name      string
	cpf       vo.CPF
	secret    string
	balance   int
	createdAt time.Time
}

var (
	_           account.Account = _account{}
	_nilAccount                 = _account{}
)

func NewAccount(
	name string,
	cpf string,
	secret string,
) (_account, error) {
	if len(name) < _nameMinLength {
		return _nilAccount, account.ErrInvalidAccountName
	}

	if len(secret) < _secretMinSize {
		return _nilAccount, account.ErrInvalidAccountSecret
	}

	_cpf, err := vo.ParseCPF(cpf)
	if err != nil {
		return _nilAccount, err
	}

	id := vo.NewAccountID()

	return _account{
		id:        id,
		name:      name,
		cpf:       _cpf,
		secret:    secret,
		balance:   _initialBalance,
		createdAt: time.Now(),
	}, nil
}

func (a _account) ID() string {
	return a.id.Value()
}

func (a _account) Name() string {
	return a.name
}

func (a _account) CPF() string {
	return a.cpf.Value()
}

func (a _account) Secret() string {
	return a.secret
}

func (a _account) Balance() int {
	return a.balance
}

func (a _account) CreatedAt() time.Time {
	return a.createdAt
}
