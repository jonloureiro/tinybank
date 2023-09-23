package accounts

import (
	"fmt"
	"time"

	"github.com/jonloureiro/tiny-bank/internal"
	"github.com/jonloureiro/tiny-bank/internal/accounts/vos"
)

const (
	_nameMinLength  = 3
	_secretMinSize  = 6
	_initialBalance = 100_00
)

type accountEntity struct {
	id        vos.AccountID
	name      string
	cpf       vos.CPF
	secret    string
	balance   int
	createdAt time.Time
}

var NilAccountEntity = accountEntity{}

func New(
	name string,
	cpf string,
	secret string,
) (accountEntity, error) {
	if len(name) < _nameMinLength {
		return NilAccountEntity, fmt.Errorf(
			"%w:invalid name", internal.ErrFailedDependency,
		)
	}

	if len(secret) < _secretMinSize {
		return NilAccountEntity, fmt.Errorf(
			"%w:invalid secret", internal.ErrFailedDependency,
		)
	}

	id := vos.NewAccountID()

	cpfVO, err := vos.NewCPF(cpf)
	if err != nil {
		return NilAccountEntity, err
	}

	return accountEntity{
		id:        id,
		name:      name,
		cpf:       cpfVO,
		secret:    secret,
		balance:   _initialBalance,
		createdAt: time.Now(),
	}, nil
}

func (a accountEntity) ID() string {
	return a.id.Value()
}

func (a accountEntity) Name() string {
	return a.name
}

func (a accountEntity) CPF() string {
	return a.cpf.Value()
}

func (a accountEntity) Secret() string {
	return a.secret
}

func (a accountEntity) Balance() int {
	return a.balance
}

func (a accountEntity) CreatedAt() time.Time {
	return a.createdAt
}
