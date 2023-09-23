package vos

import (
	"fmt"

	"github.com/jonloureiro/tiny-bank/internal/common"
	"github.com/jonloureiro/tiny-bank/pkg/stringid"
)

type AccountID struct {
	value stringid.StringID
}

var nilAccountID = AccountID{}

func NewAccountID() AccountID {
	return AccountID{stringid.New()}
}

func ParseAccountID(s string) (AccountID, error) {
	id, err := stringid.FromString(s)
	if err != nil {
		return nilAccountID, fmt.Errorf(
			"%w:%w", common.ErrFailedDependency, err,
		)
	}

	return AccountID{id}, nil
}

func (a AccountID) Value() string {
	return a.value.String()
}
