package vo_test

import (
	"testing"

	"github.com/jonloureiro/tinybank/legacy/app/vo"
)

func TestCPF(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		want := "69029890100"
		cpf, _ := vo.NewCPF(want)
		got := cpf.Value()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("validate", func(t *testing.T) {
		testCases := map[string]struct {
			expectedErr error
			cpf         string
		}{
			"invalid length": {
				expectedErr: vo.ErrInvalidLength,
				cpf:         "1",
			},
			// TODO: create more test cases
		}
		for desc, tC := range testCases {
			t.Run(desc, func(t *testing.T) {
				_, err := vo.NewCPF(tC.cpf)
				if err != tC.expectedErr {
					t.Errorf("got: %v, expected error: %v", err, tC.expectedErr)
				}
			})
		}
	})
}
