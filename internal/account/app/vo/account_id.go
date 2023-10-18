package vo

import (
	"github.com/jonloureiro/tinybank/internal/account"
	"github.com/jonloureiro/tinybank/pkg/strid"
)

type AccountID struct {
	value strid.StringID
}

var nilAccountID = AccountID{}

func NewAccountID() AccountID {
	return AccountID{strid.New()}
}

func ParseAccountID(s string) (AccountID, error) {
	id, err := strid.FromString(s)
	if err != nil {
		return nilAccountID, account.ErrInvalidAccountID
	}

	return AccountID{id}, nil
}

func (a AccountID) Value() string {
	return a.value.String()
}
