package app

import (
	"fmt"
	"time"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/app/vos"
	"github.com/jonloureiro/tiny-bank/internal/common"
)

const (
	_nameMinLength  = 3
	_secretMinSize  = 6
	_initialBalance = 100_00
)

type _account struct {
	id        vos.AccountID
	name      string
	cpf       vos.CPF
	secret    string
	balance   int
	createdAt time.Time
}

var (
	_           accounts.Account = _account{}
	_nilAccount                  = _account{}
)

func NewAccount(
	name string,
	cpf string,
	secret string,
) (_account, error) {
	if len(name) < _nameMinLength {
		return _nilAccount, fmt.Errorf(
			"%w:invalid name", common.ErrFailedDependency,
		)
	}

	if len(secret) < _secretMinSize {
		return _nilAccount, fmt.Errorf(
			"%w:invalid secret", common.ErrFailedDependency,
		)
	}

	id := vos.NewAccountID()

	cpfVO, err := vos.NewCPF(cpf)
	if err != nil {
		return _nilAccount, err
	}

	return _account{
		id:        id,
		name:      name,
		cpf:       cpfVO,
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
