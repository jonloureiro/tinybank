package id

import (
	"github.com/google/uuid"
)

type ID string

// New returns a new ID.
func New() ID {
	return ID(uuid.NewString())
}
