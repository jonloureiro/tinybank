package strid

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidStringID = errors.New("invalid string id")

type StringID struct {
	value uuid.UUID
}

func New() StringID {
	return StringID{uuid.New()}
}

func FromString(s string) (StringID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return StringID{}, ErrInvalidStringID
	}

	return StringID{id}, nil
}

func (s StringID) String() string {
	return s.value.String()
}
