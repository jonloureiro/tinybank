package id_test

import (
	"testing"

	"github.com/jonloureiro/go-challenge/extensions/id"
)

func TestNew(t *testing.T) {
	ID1 := id.New()
	ID2 := id.New()

	if ID1 == ID2 {
		t.Errorf("expected ID1 to be different from ID2")
	}
}
