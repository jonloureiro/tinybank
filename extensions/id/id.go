package id

import (
	"github.com/google/uuid"
)

// New returns a new ID.
func New() string {
	return uuid.NewString()
}
